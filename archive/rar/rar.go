package rar

import (
	"os"

	"github.com/baulk/baulkarc/archive/basics"
	"github.com/nwaples/rardecode"
)

// Extractor type
type Extractor struct {
	fd *os.File
	rr *rardecode.Reader     // underlying stream reader
	rc *rardecode.ReadCloser // supports multi-volume archives (files only)
	es *basics.ExtractSetting
}

// Matched Magic
func Matched(buf []byte) bool {
	return len(buf) > 6 &&
		buf[0] == 0x52 && buf[1] == 0x61 && buf[2] == 0x72 &&
		buf[3] == 0x21 && buf[4] == 0x1A && buf[5] == 0x7 &&
		(buf[6] == 0x0 || buf[6] == 0x1)
}

// NewExtractor new extractor
func NewExtractor(fd *os.File, es *basics.ExtractSetting) (*Extractor, error) {
	rr, err := rardecode.NewReader(fd, es.Password)
	if err != nil {
		fd.Close()
		return nil, err
	}
	return &Extractor{rr: rr}, nil
}

// Close fd
func (e *Extractor) Close() error {
	return e.fd.Close()
}

// Extract file
func (e *Extractor) Extract(destination string) error {
	//e.rr.Next()
	return nil
}
