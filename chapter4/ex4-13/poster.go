package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

/*
The JSON-based web service of the Open Movie Database lets you search https://omdbapi.com/
for a movie by name and download its poster image. Write a tool poster that
downloads the poster image for the movie named on the command line.
*/

type MovieResp struct {
	Title    string
	Year     string
	Poster   string
	Response string
	Error    string
}

func main() {
	key, ok := os.LookupEnv("API_KEY")
	if !ok {
		log.Fatalf("Environment variable API_KEY not found")
	}

	if len(os.Args) < 2 {
		log.Fatal("Usage: poster <movie>")
	}

	u := url.URL{
		Scheme: "https",
		Host:   "www.omdbapi.com",
		RawQuery: url.Values{
			"t":      os.Args[1:],
			"apikey": []string{key},
		}.Encode(),
	}

	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var movie MovieResp
	if err = json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		log.Fatal(err)
	}

	if movie.Response != "True" {
		log.Fatal(movie.Error)
	}

	fmt.Printf("poster: %v\n", movie.Poster)

	resp, err = http.Get(movie.Poster)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	f, err := os.Create("poster.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	chunk := make([]byte, 4)
	for {
		_, err := resp.Body.Read(chunk)
		if err == io.EOF {
			break
		}
		_, err = f.Write(chunk)
		if err != nil {
			log.Fatal(err)
		}

	}
}
