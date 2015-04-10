package main

import (
	"os"
	"path/filepath"
	"github.com/strayard/strayard/pkg/cmd/strayard"
)

func main() {
	basename := filepath.Base(os.Args[0])
	command := strayard.CommandFor(basename)
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}