// Package client ...
package client

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"net"
)

// Client ...
type Client struct {
	url, hash, password string
	net.Conn
}

// New ...
func New(url, password string) *Client {
	return &Client{url, "", password, nil}
}

// Connect ...
func (c *Client) Connect() error {
	con, conErr := net.Dial("tcp", c.url)
	if conErr != nil {
		return fmt.Errorf("error during connecting to sever on url %v error value %+v", c.url, conErr)
	}
	err := gob.NewEncoder(con).Encode(c.password)
	if err != nil {
		return fmt.Errorf("error during sending password to server with value %v", err)
	}
	var pass = ""
	err = gob.NewDecoder(con).Decode(&pass)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	c.hash = pass
	return nil
}
