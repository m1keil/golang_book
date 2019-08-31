// Build a tool that lets users create, read, update, and close GitHub issues from the command line, invoking their
// preferred text editor when substantial text input is required.
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const API = "https://api.github.com"
const USAGE = `A simple github command line utility to allow create/read/update and close Github issues.
Credentails must be saved in GH_USER & GH_PASS.
Will use whatever editor is defined in EDITOR environment variable (defaults to vi)`

func main() {
	var repo, action string
	var number uint

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%v\n\n", USAGE)
		fmt.Fprintln(flag.CommandLine.Output(), "Flags:")
		flag.PrintDefaults()
	}
	flag.StringVar(&action, "action", "read", "`type` is one of: \"create\", \"read\", \"update\" or \"close\"")
	flag.StringVar(&repo, "repo", "", "repository `slug` (\"owner/repo\")")
	flag.UintVar(&number, "issue", 0, "issue `number`. can be omitted when action is \"create\"")
	flag.Parse()

	if repo == "" {
		flag.Usage()
		os.Exit(1)
	}
	if action != "create" && number == 0 {
		flag.Usage()
		os.Exit(1)
	}

	var issue Issue
	var err error
	switch action {
	case "create":
		err = edit(&issue)
		if err != nil {
			break
		}
		issue, err = create(repo, issue)
	case "read":
		issue, err = read(repo, int(number))
	case "update":
		issue, err = read(repo, int(number))
		if err != nil {
			break
		}
		err = edit(&issue)
		if err != nil {
			break
		}
		issue, err = update(repo, issue)
	case "close":
		issue, err = clse(repo, int(number))
	default:
		flag.Usage()
		os.Exit(1)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Printf("Title: %v\n", issue.Title)
	fmt.Printf("Number: %v\n", issue.Number)
	fmt.Printf("Status: %v\n", issue.State)
	fmt.Println("------")
	fmt.Println(issue.Body)
}

type Issue struct {
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
	State  string `json:"state,omitempty"`
	Number int    `json:"number,omitempty"`
}

type GhError struct {
	Message string `json:"message"`
}

func create(repo string, issue Issue) (Issue, error) {
	body, err := json.Marshal(issue)
	if err != nil {
		return Issue{}, err
	}

	resp, err := request(API+"/repos/"+repo+"/issues", "POST", bytes.NewReader(body))
	if err != nil {
		return Issue{}, err
	}

	var rissue Issue
	if err := json.Unmarshal(resp, &rissue); err != nil {
		return Issue{}, err
	}

	return rissue, nil
}

func read(repo string, id int) (Issue, error) {
	resp, err := request(API+"/repos/"+repo+"/issues/"+strconv.Itoa(id), "GET", nil)
	if err != nil {
		return Issue{}, err
	}

	var rissue Issue
	if err := json.Unmarshal(resp, &rissue); err != nil {
		return Issue{}, err
	}

	return rissue, nil
}

func update(repo string, issue Issue) (Issue, error) {
	body, err := json.Marshal(issue)
	if err != nil {
		return Issue{}, err
	}

	resp, err := request(API+"/repos/"+repo+"/issues/"+strconv.Itoa(issue.Number), "PATCH", bytes.NewReader(body))
	if err != nil {
		return Issue{}, err
	}

	var rissue Issue
	if err := json.Unmarshal(resp, &rissue); err != nil {
		return Issue{}, err
	}

	return rissue, nil
}

func clse(repo string, id int) (Issue, error) {
	body, err := json.Marshal(Issue{State: "closed"})
	if err != nil {
		return Issue{}, err
	}

	resp, err := request(API+"/repos/"+repo+"/issues/"+strconv.Itoa(id), "PATCH", bytes.NewReader(body))
	if err != nil {
		return Issue{}, err
	}

	var rissue Issue
	if err := json.Unmarshal(resp, &rissue); err != nil {
		return Issue{}, err
	}

	return rissue, nil
}

func edit(issue *Issue) error {
	// create unique temp file
	f, _ := ioutil.TempFile("", "ghissue_*")

	// if existing issue is not blank, need to pre-fill the tmp file
	if issue.Title != "" {
		_, err := f.WriteString(issue.Title + "\n" + issue.Body)
		if err != nil {
			return err
		}
	}

	f.Close()

	// figure out which editor to use (try vi as default)
	editor, ok := os.LookupEnv("EDITOR")
	if !ok || editor == "" {
		_, err := os.Stat("/usr/bin/vi")
		if err != nil {
			return err
		}
		editor = "/usr/bin/vi"
	}

	// exec editor and point to the temp file
	cmd := &exec.Cmd{
		Path:   editor,
		Args:   []string{editor, f.Name()},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	if err := cmd.Run(); err != nil {
		return err
	}

	var content []string
	// read content of the file
	fh, err := os.Open(f.Name())
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	// remove tmp file
	os.Remove(f.Name()) // ignore error

	if len(content) == 0 {
		return fmt.Errorf("canceled")
	}
	issue.Title = content[0]
	issue.Body = strings.Join(content[1:], "\n")

	return nil
}

func request(url, method string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(os.Getenv("GH_USER"), os.Getenv("GH_PASS"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 && resp.StatusCode < 500 {
		var gherror GhError
		if err := json.NewDecoder(resp.Body).Decode(&gherror); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("error: %v", gherror.Message)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}
