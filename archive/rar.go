package archive

import "github.com/nwaples/rardecode"

// Rar type
type Rar struct {
	rr *rardecode.Reader     // underlying stream reader
	rc *rardecode.ReadCloser // supports multi-volume archives (files only)
}
