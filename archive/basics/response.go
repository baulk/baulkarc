package basics

import (
	"encoding/json"
	"os"
	"path"
	"path/filepath"
)

// File compress file
type File struct {
	Path        string `json:"path"`
	Destination string `json:"destination"`
	Name        string `json:"name,omitempty"`        // if not exists use filepath.Base
	Executabled bool   `json:"executabled,omitempty"` // when mark executabled. a script create under windows can run linux
}

// BuildPath todo
func (file *File) BuildPath() string {
	destination := filepath.ToSlash(file.Destination)
	if len(file.Name) == 0 {
		return path.Join(destination, filepath.Base(file.Path))
	}
	return path.Join(destination, file.Name)
}

// ResponseFile todo
type ResponseFile struct {
	cwd           string
	Destination   string   `json:"destination"`
	CompressLevel int      `json:"level,omitempty"`
	Method        string   `json:"method,omitempty"`
	Files         []File   `json:"files,omitempty"`
	Dirs          []string `json:"dirs,omitempty"`
}

// DefaultResponseFile default response file
const DefaultResponseFile = "compress.rsp.json"

// NewResponseFile response file
func NewResponseFile(src string) (*ResponseFile, error) {
	fd, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	rsp := ResponseFile{cwd: filepath.Dir(src)}
	if err := json.NewDecoder(fd).Decode(&rsp); err != nil {
		return nil, err
	}
	return &rsp, nil
}
