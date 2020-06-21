package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/baulk/baulkarc/utilities"
)

// IsDebugMode todo
var IsDebugMode bool = false

// DbgPrint todo
func DbgPrint(format string, v ...interface{}) {
	if IsDebugMode {
		s := fmt.Sprintf(format, v...)
		strings.TrimSuffix(s, "\n")
		utilities.WriteFile(os.Stderr, "\x1b[33m", s, "\x1b[0m")
	}
}

func main() {
}
