package archive

import "github.com/klauspost/compress/zstd"

// tar

type decompress struct {
	r zstd.Decoder
}
