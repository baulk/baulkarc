package zip

import (
	"archive/zip"
	"compress/bzip2"
	"compress/flate"
	"errors"
	"io"
	"io/ioutil"
	"sync"

	"github.com/klauspost/compress/zstd"
	"github.com/ulikunitz/xz"
	"github.com/ulikunitz/xz/lzma"
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

func newZstdReader(r io.Reader) io.ReadCloser {
	zr, err := zstd.NewReader(r)
	if err != nil {
		return nil
	}
	return zr.IOReadCloser()
}

func zipRegisterDecompressor() {
	zip.RegisterDecompressor(uint16(BZIP2), newBzip2Reader)
	zip.RegisterDecompressor(uint16(LZMA), newLzmaReader)
	zip.RegisterDecompressor(uint16(XZ), newXzReader)
	zip.RegisterDecompressor(uint16(ZSTD), newZstdReader)
	zip.RegisterDecompressor(uint16(WINZIPZSTD), newZstdReader)
}
