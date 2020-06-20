package s7z

import (
	"io"
	"os"

	"github.com/baulk/baulkarc/archive/settings"
	"github.com/baulk/baulkarc/go7z"
)

// Extractor type
type Extractor struct {
	fd  *os.File
	szr *go7z.Reader
	es  *settings.ExtractSetting
}

// Matched magic
func Matched(buf []byte) bool {
	return len(buf) > 5 &&
		buf[0] == 0x37 && buf[1] == 0x7A && buf[2] == 0xBC &&
		buf[3] == 0xAF && buf[4] == 0x27 && buf[5] == 0x1C
}

//NewExtractor new tar extractor
func NewExtractor(fd *os.File, es *settings.ExtractSetting) (*Extractor, error) {
	st, err := fd.Stat()
	if err != nil {
		fd.Close()
		return nil, err
	}
	r, err := go7z.NewReader(fd, st.Size())
	if err != nil {
		fd.Close()
		return nil, err
	}
	e := &Extractor{szr: r, fd: fd, es: es}
	e.szr.Options.SetPassword(es.Password)
	e.szr.Options.SetPasswordCallback(es.PassworldCallback)
	return e, nil
}

// Close fd
func (e *Extractor) Close() error {
	return e.fd.Close()
}

// Extract file
func (e *Extractor) Extract(destination string) error {
	for {
		hdr, err := e.szr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if hdr.IsEmptyStream && !hdr.IsEmptyFile {
			if err := os.MkdirAll(hdr.Name, os.ModePerm); err != nil {
				return err
			}
			continue
		}
		f, err := os.Create(hdr.Name)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := io.Copy(f, e.szr); err != nil {
			return err
		}
	}
	return nil
}
