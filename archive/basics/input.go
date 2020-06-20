package basics

// File compress file
type File struct {
	Path        string `json:"path"`
	Destination string `json:"destination"`
	Name        string `json:"name,omitempty"`        // if not exists use filepath.Base
	Executabled bool   `json:"executabled,omitempty"` // when mark executabled. a script create under windows can run linux
}

// CompressionInput todo
type CompressionInput struct {
	Destination string   `json:"destination"`
	Method      string   `json:"method,omitempty"`
	Files       []File   `json:"files,omitempty"`
	Dirs        []string `json:"dirs,omitempty"`
}
