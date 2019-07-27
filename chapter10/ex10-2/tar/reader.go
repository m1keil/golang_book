package tar

import (
	"archive/tar"
	"golang/chapter10/ex10-2/archive"
	"io"
	"os"
)

func init() {
	archive.RegisterFormat("tar", "\x75\x73\x74\x61\x72\x20\x20\x00", 0x101, Extract)
}

func Extract(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	tr := tar.NewReader(file)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return err
		}

		dst, err := os.OpenFile(hdr.Name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.FileMode(hdr.Mode))
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err := io.Copy(dst, tr); err != nil {
			return err
		}
	}

	return nil
}

