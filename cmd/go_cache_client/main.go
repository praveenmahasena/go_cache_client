// Package main ...
package main

import (
	"fmt"
	"os"

	"github.com/praveenmahasena/gocacheclient/internal"
)

func main() {
	if err := internal.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1) // -1 should not be here but meh
	}
}
