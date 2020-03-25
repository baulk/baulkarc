package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

// ZipDecompressor todo
type ZipDecompressor struct {
	OverwriteExisting bool
	MkdirAll          bool
	reader            *zip.Reader
	ridx              int
}

func (z *ZipDecompressor) initialize(reader io.Reader, size int64) error {
	inRdrAt, ok := reader.(io.ReaderAt)
	if !ok {
		return fmt.Errorf("reader must be io.ReaderAt")
	}
	if z.reader != nil {
		return fmt.Errorf("zip archive is already open for reading")
	}
	var err error
	if z.reader, err = zip.NewReader(inRdrAt, size); err != nil {
		return fmt.Errorf("creating reader: %v", err)
	}
	z.ridx = 0
	return nil
}

// Decompression todo
func (z *ZipDecompressor) Decompression(reader io.Reader, size int64, dest string) error {
	if !pathExists(dest) && z.MkdirAll {
		if err := os.MkdirAll(dest, 0755); err != nil {
			return fmt.Errorf("preparing destination: %v", err)
		}
	}
	if err := z.initialize(reader, size); err != nil {
		return err
	}
	files := len(z.reader.File)
	DbgPrint("%d files", files)
	return nil
}
