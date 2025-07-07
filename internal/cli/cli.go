// Package cli ...
package cli

import (
	"bufio"
	"os"
)

// Cli ...
type Cli struct {
	*os.File
}

// New ...
func New(f *os.File) *Cli {
	return &Cli{f}
}

// ReadFrom ...
func (c *Cli) ReadFrom(messageCh chan string) {
	buff := bufio.NewScanner(c)
	for buff.Scan() {
		messageCh <- buff.Text()
	}
}
