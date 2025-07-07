// Package internal ...
package internal

import (
	"github.com/praveenmahasena/gocacheclient/internal/client"
)

// Run to start app
func Run() error {
	client := client.New("localhost:42069", "pass")
	if err := client.Connect(); err != nil {
		return err
	}
	return nil
}
