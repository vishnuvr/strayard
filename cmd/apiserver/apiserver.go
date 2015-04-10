package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	s := app.NewAPIServer()

	if err := s.Run([]string{}); err != nil {
		fmt.Fprint(os.Stderr, err.Error)
		os.Exit(1)
	}
}
