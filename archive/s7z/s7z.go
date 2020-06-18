package s7z

import (
	"os"

	"github.com/baulk/baulkarc/archive/rules"
)

// Extractor type
type Extractor struct {
	fd *os.File
	es *rules.ExtractSetting
}
