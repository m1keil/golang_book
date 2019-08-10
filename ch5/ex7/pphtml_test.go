package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

var TestData = `
<html>
</html>
`

func TestMain(m *testing.M) {
	switch os.Getenv("TEST_MAIN") {
	case "crasher":
		fmt.Println("test12345")
		// os.Exit(5)
		doc, _ := html.Parse(strings.NewReader(TestData))
		forEachNode(doc, startElement, endElement)
	default:
		fmt.Println("default")
		os.Exit(m.Run())
	}
}

func TestPPHTML(t *testing.T) {
	fmt.Println("running")
	cmd := exec.Command(os.Args[0])
	cmd.Env = append(os.Environ(), "TEST_MAIN=crasher")
	if err := cmd.Run(); err != nil {
		t.Fatalf("process err %v, want exit status 0", err)
	}

	fmt.Printf("output: %v %v\n", cmd.Stdout, cmd.ProcessState)

}
