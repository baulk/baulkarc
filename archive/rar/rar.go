package rar

import (
	"io"
	"os"

	"github.com/baulk/baulkarc/archive/rules"
	"github.com/nwaples/rardecode"
)

// Extractor type
type Extractor struct {
	fd *os.File
	rr *rardecode.Reader     // underlying stream reader
	rc *rardecode.ReadCloser // supports multi-volume archives (files only)
	es *rules.ExtractSetting
}

// Matched Magic
func Matched(buf []byte) bool {
	return len(buf) > 6 &&
		buf[0] == 0x52 && buf[1] == 0x61 && buf[2] == 0x72 &&
		buf[3] == 0x21 && buf[4] == 0x1A && buf[5] == 0x7 &&
		(buf[6] == 0x0 || buf[6] == 0x1)
}

// NewExtractor new extractor
func NewExtractor(fd *os.File, size int64, es *rules.ExtractSetting) (*Extractor, error) {
	if _, err := fd.Seek(0, io.SeekStart); err != nil {
		fd.Close()
		return nil, err
	}
	rr, err := rardecode.NewReader(fd, "")
	if err != nil {
		return nil, err
	}
	e := &Extractor{rr: rr}
	return e, nil
}
