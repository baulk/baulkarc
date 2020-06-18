package s7z

import (
	"os"

	"github.com/baulk/baulkarc/archive/rules"
	"github.com/baulk/baulkarc/go7z"
)

// Extractor type
type Extractor struct {
	fd  *os.File
	s7r *go7z.Reader
	es  *rules.ExtractSetting
}

// Matched magic
func Matched(buf []byte) bool {
	return len(buf) > 5 &&
		buf[0] == 0x37 && buf[1] == 0x7A && buf[2] == 0xBC &&
		buf[3] == 0xAF && buf[4] == 0x27 && buf[5] == 0x1C
}
