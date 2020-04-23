package main

import "io"

// Decompressor todo
type Decompressor interface {
	Decompression(reader io.Reader, size int64, dest string) error
}
