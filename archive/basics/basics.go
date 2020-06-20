package basics

// ExtractSetting todo
type ExtractSetting struct {
	OverwriteExisting bool
	MkdirAll          bool
	Encoding          string
	Password          string
	PassworldCallback func() string
}

// Compression Algorithm
const (
	None = iota
	Deflate
	GZ
	BZip2
	LZMA
	XZ
	LZ4
	Brotli
	Zstandard
)
