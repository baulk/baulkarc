package zip

import (
	"archive/zip"
	"os"

	"github.com/baulk/baulkarc/archive/basics"
	"golang.org/x/text/encoding"
)

// CompressionMethod compress method see https://pkware.cachefly.net/webdocs/casestudies/APPNOTE.TXT
type CompressionMethod uint16

// CompressionMethod
// value
const (
	Store      CompressionMethod = 0
	Deflate    CompressionMethod = 8
	BZIP2      CompressionMethod = 12
	LZMA       CompressionMethod = 14
	ZSTD       CompressionMethod = 20
	LZMA2      CompressionMethod = 33
	WINZIPZSTD CompressionMethod = 93
	XZ         CompressionMethod = 95
	JPEG       CompressionMethod = 96
	WavPack    CompressionMethod = 97
	PPMd       CompressionMethod = 98
	AES        CompressionMethod = 99
)

func init() {
	zipRegisterDecompressor()
	zipRegisterCompressor()
}

// Matched magic
func Matched(buf []byte) bool {
	return (len(buf) > 3 && buf[0] == 0x50 && buf[1] == 0x4B &&
		(buf[2] == 0x3 || buf[2] == 0x5 || buf[2] == 0x7) &&
		(buf[3] == 0x4 || buf[3] == 0x6 || buf[3] == 0x8))
}

// Extractor todo
type Extractor struct {
	fd  *os.File
	zr  *zip.Reader
	dec *encoding.Decoder
	es  *basics.ExtractSetting
}

// NewExtractor new extractor
func NewExtractor(fd *os.File, es *basics.ExtractSetting) (*Extractor, error) {
	st, err := fd.Stat()
	if err != nil {
		fd.Close()
		return nil, err
	}
	zr, err := zip.NewReader(fd, st.Size())
	if err != nil {
		fd.Close()
		return nil, err
	}
	e := &Extractor{fd: fd, zr: zr, es: es}
	if ens := os.Getenv("ZIP_ENCODING"); len(ens) != 0 {
		e.initializeEncoder(ens)
	}
	return e, nil
}

// Close fd
func (e *Extractor) Close() error {
	return e.fd.Close()
}

// Extract file
func (e *Extractor) Extract(destination string) error {

	return nil
}
