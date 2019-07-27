package archive

import (
	"fmt"
	"os"
	"sync"
)

type format struct {
	id string
	magic string
	offset int64
	extractAll func(string) error
}

var formats []format
var lock sync.Mutex

func RegisterFormat(id, magic string, offset int64, extract func(string) error) {
	lock.Lock()
	formats = append(formats, format{id, magic, offset, extract})
	lock.Unlock()
}

func sniff(file *os.File) format {
	for _, f := range formats {
		magic := make([]byte, len(f.magic))
		_, err := file.ReadAt(magic, f.offset)
		if err == nil && f.magic == string(magic) {
			return f
		}
	}
	return format{}
}


func ExtractArchive(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	format :=  sniff(f)
	f.Close()

	if format.extractAll == nil {
		return fmt.Errorf("no format found")
	}

	return format.extractAll(path)

}