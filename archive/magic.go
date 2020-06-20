package archive

import (
	"errors"
	"io"
	"os"

	"github.com/baulk/baulkarc/archive/rules"
	"github.com/baulk/baulkarc/archive/tar"
	"github.com/baulk/baulkarc/archive/zip"
)

func readMagic(fd *os.File) ([]byte, error) {
	buf := make([]byte, 0, 512)
	l, err := fd.Read(buf)
	if err != nil {
		return nil, err
	}
	if _, err := fd.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}
	return buf[0:l], nil
}

// NewExtractor todo
func NewExtractor(file string, es *rules.ExtractSetting) (Extractor, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	mb, err := readMagic(fd)
	if err != nil {
		return nil, err
	}
	if zip.Matched(mb) {
		e, err := zip.NewExtractor(fd, es)
		if err != nil {
			return nil, err
		}
		return e, nil
	}
	if tar.Matched(mb) {
		e, _ := tar.NewExtractor(fd, es)
		return e, nil
	}
	if al := tar.MatchExtension(file); al != rules.None {
		e, err := tar.NewBrewingExtractor(fd, es, al)
		if err != nil {
			return nil, err
		}
		return e, nil
	}
	return nil, errors.New("unsupport format")
}
