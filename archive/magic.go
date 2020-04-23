package archive

import (
	"io"
	"os"
)

// NewDecompressorFromFile todo
func NewDecompressorFromFile(file string) (*Decompressor, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, 0, 256)
	l, err := fd.Read(buf)
	if err != nil {
		_ = fd.Close()
		return nil, err
	}
	// check magic
	if l == 1 {
	}
	if _, err = fd.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}
	return nil, nil
}
