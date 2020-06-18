package tar

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"github.com/andybalholm/brotli"
	"github.com/baulk/baulkarc/archive/rules"
	"github.com/dsnet/compress/bzip2"
	"github.com/klauspost/compress/zstd"
	"github.com/pierrec/lz4/v3"
	"github.com/ulikunitz/xz"
)

// tar

// Extractor type
type Extractor struct {
	fd *os.File
	r  *tar.Reader
	es *rules.ExtractSetting
}

// Close fd
func (e *Extractor) Close() error {
	return e.fd.Close()
}

// Extract file
func (e *Extractor) Extract(destination string) error {
	for {
		hdr, err := e.r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if hdr.Size != 0 {

		}
	}
	return nil
}

// BrewingExtractor todo
type BrewingExtractor struct {
	fd  *os.File
	r   *tar.Reader
	mwr io.ReadCloser
	es  *rules.ExtractSetting
}

// Matched todo
func Matched(buf []byte) bool {
	return len(buf) > 261 &&
		buf[257] == 0x75 && buf[258] == 0x73 &&
		buf[259] == 0x74 && buf[260] == 0x61 &&
		buf[261] == 0x72
}

// MatchExtension todo
func MatchExtension(name string) int {
	if runtime.GOOS == "windows" {
		name = strings.ToLower(name)
	}
	if strings.HasSuffix(name, ".tar.gz") || strings.HasSuffix(name, ".tgz") {
		return rules.GZ
	}
	if strings.HasSuffix(name, ".tar.bz2") || strings.HasSuffix(name, ".tbz2") {
		return rules.BZip2
	}
	if strings.HasSuffix(name, ".tar.br") || strings.HasSuffix(name, ".tbr") {
		return rules.Brotli
	}
	if strings.HasSuffix(name, ".tar.zst") {
		return rules.Zstandard
	}
	if strings.HasSuffix(name, ".tar.xz") || strings.HasSuffix(name, ".txz") {
		return rules.XZ
	}
	if strings.HasSuffix(name, ".tar.lz4") || strings.HasSuffix(name, ".tlz4") {
		return rules.XZ
	}
	return rules.None
}

// NewBrewingExtractor todo
func NewBrewingExtractor(fd *os.File, es *rules.ExtractSetting, alg int) (*BrewingExtractor, error) {
	var err error
	e := &BrewingExtractor{es: es}
	switch alg {
	case rules.GZ:
		e.mwr, err = gzip.NewReader(fd)
		if err != nil {
			fd.Close()
			return nil, err
		}
	case rules.LZ4:
		e.mwr = ioutil.NopCloser(lz4.NewReader(fd))
	case rules.Brotli:
		e.mwr = ioutil.NopCloser(brotli.NewReader(fd))
	case rules.BZip2:
		e.mwr, err = bzip2.NewReader(fd, nil)
		if err != nil {
			fd.Close()
			return nil, err
		}
	case rules.XZ:
		r, err := xz.NewReader(fd)
		if err != nil {
			fd.Close()
			return nil, err
		}
		e.mwr = ioutil.NopCloser(r)
	case rules.Zstandard:
		dec, err := zstd.NewReader(fd)
		if err != nil {
			fd.Close()
			return nil, err
		}
		e.mwr = dec.IOReadCloser()
	default:
		fd.Close()
		return nil, fmt.Errorf("unsupport compress algorithm %d", alg)
	}
	e.fd = fd
	e.r = tar.NewReader(e.mwr)
	return e, nil
}

// Close fd
func (e *BrewingExtractor) Close() error {
	_ = e.mwr.Close()
	return e.fd.Close()
}

// Extract file
func (e *BrewingExtractor) Extract(destination string) error {

	return nil
}
