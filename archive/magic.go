package archive

import (
	"errors"
	"os"

	"github.com/baulk/baulkarc/archive/rules"
	"github.com/baulk/baulkarc/archive/zip"
)

// NewExtractor todo
func NewExtractor(file string, es *rules.ExtractSetting) (Extractor, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	sz, err := fd.Stat()
	if err != nil {
		fd.Close()
		return nil, err
	}
	buf := make([]byte, 0, 256)
	l, err := fd.Read(buf)
	if err != nil {
		_ = fd.Close()
		return nil, err
	}
	if zip.Matched(buf[0:l]) {
		e, err := zip.NewExtractor(fd, sz.Size(), es)
		if err != nil {
			return nil, err
		}
		return e, nil
	}
	return nil, errors.New("unsupport format")
}
