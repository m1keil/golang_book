package library

import "time"

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func Length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type StatefulSort struct {
	T     []*Track
	Order []func(x, y *Track) (bool, bool)
}

func (x StatefulSort) Len() int      { return len(x.T) }
func (x StatefulSort) Swap(i, j int) { x.T[i], x.T[j] = x.T[j], x.T[i] }
func (x StatefulSort) Less(i, j int) bool {
	for _, f := range x.Order {
		swap, result := f(x.T[i], x.T[j])
		if swap {
			return result
		}
	}
	return false
}

func SortByTitle(x, y *Track) (swap, result bool) {
	if x.Title != y.Title {
		return true, x.Title < y.Title
	}
	return
}

func SortByArtist(x, y *Track) (swap, result bool) {
	if x.Artist != y.Artist {
		return true, x.Artist < y.Artist
	}
	return
}

func SortByYear(x, y *Track) (swap, result bool) {
	if x.Year != y.Year {
		return true, x.Year < y.Year
	}
	return
}

func SortByLength(x, y *Track) (swap, result bool) {
	if x.Length != y.Length {
		return true, x.Length < y.Length
	}
	return
}

func SortByAlbum(x, y *Track) (swap, result bool) {
	if x.Album != y.Album {
		return true, x.Album < y.Album
	}
	return
}
