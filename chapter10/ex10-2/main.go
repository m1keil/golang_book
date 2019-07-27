package main

/*
Define a generic archive file-reading function capable of reading ZIP files (archive/zip) and POSIX tar files
(archive/tar). Use a registration mechanism similar to the one described above so that support for each file format can
be plugged in using blank imports.
*/

import (
	"fmt"
	"golang/chapter10/ex10-2/archive"
	"os"
	_ "golang/chapter10/ex10-2/zip"
	_ "golang/chapter10/ex10-2/tar"
)

func main() {
	err := archive.ExtractArchive(os.Args[1]); if err != nil {
		fmt.Println("error:", err)
	}
}

