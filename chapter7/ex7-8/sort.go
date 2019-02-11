package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

/*
  Many GUIs provide a table widget with a stateful multi-tier sort: the primary
  sort key is the most recently clicked column head, the secondary sort key is
  the second-most recently clicked column head, and so on. Define an
  implementation of sort.Interface for use by such a table. Compare that
  approach with repeated sorting using sort.Stable.
*/

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type statefulSort struct {
	t     []*Track
	order []func(x, y *Track) (bool, bool)
}

func (x statefulSort) Len() int      { return len(x.t) }
func (x statefulSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }
func (x statefulSort) Less(i, j int) bool {
	for _, f := range x.order {
		swap, result := f(x.t[i], x.t[j])
		if swap {
			return result
		}
	}
	return false
}

func sortByTitle(x, y *Track) (swap, result bool) {
	if x.Title != y.Title {
		return true, x.Title < y.Title
	}
	return
}

func sortByArtist(x, y *Track) (swap, result bool) {
	if x.Artist != y.Artist {
		return true, x.Artist < y.Artist
	}
	return
}

func sortByYear(x, y *Track) (swap, result bool) {
	if x.Year != y.Year {
		return true, x.Year < y.Year
	}
	return
}

func printTracks(tracks []*Track) {
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
	sort.Sort(statefulSort{tracks,
		[]func(x, y *Track) (bool, bool){
			sortByTitle,
			sortByYear,
			// sortByArtist,
		},
	})
	printTracks(tracks)

}
