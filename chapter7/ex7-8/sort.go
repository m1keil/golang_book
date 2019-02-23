package main

import (
	"fmt"
	"golang/chapter7/ex7-8/library"
	"os"
	"sort"
	"text/tabwriter"
)

/*
  Many GUIs provide a table widget with a stateful multi-tier sort: the primary
  sort key is the most recently clicked column head, the secondary sort key is
  the second-most recently clicked column head, and so on. Define an
  implementation of sort.Interface for use by such a table. Compare that
  approach with repeated sorting using sort.Stable.
*/

var tracks = []*library.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, library.Length("3m38s")},
	{"Go", "Moby", "Moby", 1992, library.Length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, library.Length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, library.Length("4m24s")},
}

func printTracks(tracks []*library.Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

func main() {
	sort.Sort(library.StatefulSort{tracks,
		[]func(x, y *library.Track) (bool, bool){
			library.SortByTitle,
			library.SortByYear,
			// sortByArtist,
		},
	})
	printTracks(tracks)

}
