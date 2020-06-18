package tar

import (
	"archive/tar"
	"io"
	"os"

	"github.com/baulk/baulkarc/archive/rules"
)

// tar

// Extractor type
type Extractor struct {
	fd *os.File
	r  *tar.Reader
	es *rules.ExtractSetting
}

// BrewingExtractor todo
type BrewingExtractor struct {
	fd *os.File
	r  *tar.Reader
	er *io.ReadCloser
	es *rules.ExtractSetting
}

// Matched todo
func Matched(buf []byte) bool {
	return len(buf) > 261 &&
		buf[257] == 0x75 && buf[258] == 0x73 &&
		buf[259] == 0x74 && buf[260] == 0x61 &&
		buf[261] == 0x72
}

// MatchExtension todo
func MatchExtension(name string) {

}
