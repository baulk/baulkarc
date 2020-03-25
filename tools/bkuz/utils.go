package main

import (
	"fmt"
	"os"
	"strings"
)

// StrCat cat strings:
// You should know that StrCat gradually builds advantages
// only when the number of parameters is> 2.
func StrCat(sv ...string) string {
	var sb strings.Builder
	var size int
	for _, s := range sv {
		size += len(s)
	}
	sb.Grow(size)
	for _, s := range sv {
		_, _ = sb.WriteString(s)
	}
	return sb.String()
}

// IsDebugMode todo
var IsDebugMode bool = false

// DbgPrint todo
func DbgPrint(format string, v ...interface{}) {
	if IsDebugMode {
		s := fmt.Sprintf(format, v...)
		strings.TrimSuffix(s, "\n")
		os.Stderr.WriteString(StrCat("\x1b[33m", s, "\x1b[0m"))
	}
}

func pathExists(p string) bool {
	_, err := os.Stat(p)
	return !os.IsNotExist(err)
}
