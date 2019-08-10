package zip

import (
	"archive/zip"
	"golang/ch10/ex10-2/archive"
	"io"
	"log"
	"os"
)

func init() {
	archive.RegisterFormat("zip", "\x50\x4B\x03\x04", 0, Extract)
}

func Extract(path string) error {
	r, err := zip.OpenReader(path)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		dst, err := os.OpenFile(f.Name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer dst.Close()

		_, err = io.Copy(dst, rc)
		if err != nil {
			return err
		}
	}
	return nil
}