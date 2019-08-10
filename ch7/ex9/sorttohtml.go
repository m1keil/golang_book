package main

import (
	"golang/ch7/ex7-8/library"
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"
	"sync"
)

/*
 Use the html/template package (§4.6) to replace printTracks with a function
 that displays the tracks as an HTML table. Use the solution to the previous
 exercise to arrange that each click on a column head makes an HTTP request to
 sort the table.
*/

var tracks = []*library.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, library.Length("3m39s")},
	{"Go", "Delilah", "From the Roots Up", 2012, library.Length("3m38s")},
	{"Go", "Moby", "Moby", 1992, library.Length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, library.Length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, library.Length("4m24s")},
}

func generateHTML(tracks []*library.Track, w io.Writer) {
	tpl := `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Table sort</title>
	</head>
	<body>
		<table>
			<tr>
				<th><a href="/?sortBy=title">Title</a></th>
				<th><a href="/?sortBy=artist">Artist</a></th>
				<th><a href="/?sortBy=album">Album</a></th>
				<th><a href="/?sortBy=year">Year</a></th>
				<th><a href="/?sortBy=length">Length</a></th>
			</tr>
			{{range $_, $value := .}}
			<tr>
				<td>{{ $value.Title }}</td>
				<td>{{ $value.Artist }}</td>
				<td>{{ $value.Album }}</td>
				<td>{{ $value.Year }}</td>
				<td>{{ $value.Length }}</td>
			</tr>
			{{end}}
		</table>
	</body>
</html>`

	template.Must(template.New("table").Parse(tpl)).Execute(w, tracks)
}

var order []func(x, y *library.Track) (bool, bool)
var mu sync.Mutex

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	switch r.FormValue("sortBy") {
	case "title":
		order = append(order, library.SortByTitle)
	case "artist":
		order = append(order, library.SortByArtist)
	case "album":
		order = append(order, library.SortByAlbum)
	case "year":
		order = append(order, library.SortByYear)
	case "length":
		order = append(order, library.SortByLength)
	case "clear":
		order = order[:0]
	}
	mu.Unlock()

	sort.Sort(library.StatefulSort{tracks, order})

	generateHTML(tracks, w)
}
