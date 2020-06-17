package archive

import "io"

// https://ethw.org/History_of_Lossless_Data_Compression_Algorithms

// Decompressor todo
type Decompressor struct {
	in io.ReadCloser // diskfile
}

// Extractor todo
type Extractor interface {
	Extract(destination string) error
	Close() error
}

// Archiver todo
type Archiver interface {
}
