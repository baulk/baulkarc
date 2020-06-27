package main

import (
	"fmt"
	"os"

	"github.com/baulk/bkz/netutils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s url\n", os.Args[0])
		os.Exit(1)
	}
	executor := netutils.NewExecutor()
	file, err := executor.Get(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "download %s error %v\n", os.Args[1], err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "download success save to %s\n", file)
}
