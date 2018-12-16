package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

/*
The popular web comic xkcd has a JSON interface. For example, a request to https://xkcd.com/571/info.0.json
produces a detailed description of comic 571, one of many favorites. Download
each URL (once!) and build an offline index. Write a tool xkcd that, using this
index, prints the URL and transcript of each comic that matches a search term
provided on the command line.
*/

const IndexPath = "./index"

type ComicItem struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

func main() {
	files, err := ioutil.ReadDir(IndexPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		fname := filepath.Join(IndexPath, f.Name())
		content, err := ioutil.ReadFile(fname)
		if err != nil {
			log.Fatal(err)
		}

		var item ComicItem
		err = json.Unmarshal(content, &item)
		if err != nil {
			continue
		}

		term := os.Args[1]
		if strings.Contains(item.Transcript, term) {
			fmt.Printf("\nlink: https://xkcd.com/%v/\n", item.Num)
			fmt.Printf("%v\n", item.Transcript)
			fmt.Println("---------")
		}
	}
}
