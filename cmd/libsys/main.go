package main

import (
	"fmt"
	"os"

	"github.com/praveenmahasena/libsys/internal"
)

func main() {
	if err := internal.Start(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}
