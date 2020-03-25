package main

import (
	"fmt"
	"io"
)

// todo
const (
	None = iota
	Zip
	SevenZip
)

// IsZip todo
func IsZip(buf []byte) bool {
	return (len(buf) > 3 && buf[0] == 0x50 && buf[1] == 0x4B &&
		(buf[2] == 0x3 || buf[2] == 0x5 || buf[2] == 0x7) &&
		(buf[3] == 0x4 || buf[3] == 0x6 || buf[3] == 0x8))
}

// NewDecompressor new
func NewDecompressor(r io.Reader) (Decompressor, error) {
	seeker, ok := r.(io.Seeker)
	if !ok {
		return nil, fmt.Errorf("reader must be io.Seeker")
	}
	header := make([]byte, 261)
	if _, err := r.Read(header); err != nil {
		return nil, fmt.Errorf("unable read data %v", err)
	}
	var decompressor Decompressor
	if IsZip(header) {
		decompressor = &ZipDecompressor{}
	}
	// TO read data
	if _, err := seeker.Seek(0, io.SeekStart); err != nil {
		return nil, fmt.Errorf("see to begin %v", err)
	}
	return decompressor, nil
}

// // NewDecompressReader todo
// func NewDecompressReader(p string) (io.ReadCloser, int64, *Decompressor, error) {
// 	fd, err := os.Open(p)
// 	if err != nil {
// 		return nil, 0, nil, err
// 	}
// 	return fd, 0, &ZipDecompressor{}, nil
// }
