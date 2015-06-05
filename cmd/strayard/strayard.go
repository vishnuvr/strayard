package main

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/strayard/strayard/pkg/cmd/strayard"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	basename := filepath.Base(os.Args[0])
	command := strayard.CommandFor(basename)
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
