package archive

import (
	"archive/zip"
	"compress/bzip2"
	"compress/flate"
	"errors"
	"io"
	"io/ioutil"
	"sync"

	"github.com/ulikunitz/xz"
	"github.com/ulikunitz/xz/lzma"
)

// TODO support deflate64

// value
const (
	Store    uint16 = 0
	Deflate  uint16 = 8
	Deflate9 uint16 = 9
	BZIP2    uint16 = 12
	LZMA     uint16 = 14
	XZ       uint16 = 95
	JPEG     uint16 = 96
	WavPack  uint16 = 97
	PPMd     uint16 = 98
	AES      uint16 = 99
)

var bzip2ReaderPool sync.Pool

func newBzip2Reader(r io.Reader) io.ReadCloser {
	fr, ok := bzip2ReaderPool.Get().(io.ReadCloser)
	if ok {
		fr.(flate.Resetter).Reset(r, nil)
	} else {
		fr = ioutil.NopCloser(bzip2.NewReader(r))
	}
	return &pooledBzip2Reader{fr: fr}
}

type pooledBzip2Reader struct {
	mu sync.Mutex // guards Close and Read
	fr io.ReadCloser
}

func (r *pooledBzip2Reader) Read(p []byte) (n int, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.fr == nil {
		return 0, errors.New("Read after Close")
	}
	return r.fr.Read(p)
}

func (r *pooledBzip2Reader) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	var err error
	if r.fr != nil {
		err = r.fr.Close()
		bzip2ReaderPool.Put(r.fr)
		r.fr = nil
	}
	return err
}

// support xz
func newXzReader(r io.Reader) io.ReadCloser {
	xr, err := xz.NewReader(r)
	if err != nil {
		return nil
	}
	return ioutil.NopCloser(xr)
}

func newLzmaReader(r io.Reader) io.ReadCloser {
	lr, err := lzma.NewReader(r)
	if err != nil {
		return nil
	}
	return ioutil.NopCloser(lr)
}

func zipRegister() {
	zip.RegisterDecompressor(BZIP2, newBzip2Reader)
	zip.RegisterDecompressor(LZMA, newLzmaReader)
	zip.RegisterDecompressor(XZ, newXzReader)
}

func magicIsZip(buf []byte) bool {
	return (len(buf) > 3 && buf[0] == 0x50 && buf[1] == 0x4B &&
		(buf[2] == 0x3 || buf[2] == 0x5 || buf[2] == 0x7) &&
		(buf[3] == 0x4 || buf[3] == 0x6 || buf[3] == 0x8))
}
