package archive

import "io"

// https://ethw.org/History_of_Lossless_Data_Compression_Algorithms

// Decompressor todo
type Decompressor struct {
	reader io.ReadCloser
}
