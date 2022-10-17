// Package test for the application
package test

import (
	"strings"

	"github.com/ddliu/go-httpclient"
)

// Client holds the info for the HTTP client
type Client struct {
	ServerURL string
	Resp      *httpclient.Response
	JSON      map[string]interface{}
	Text      string
	Err       error
}

// HasJSON checks if content type is application/json
func (c *Client) HasJSON() bool {
	return c.Resp != nil && strings.Contains(c.Resp.Header.Get("content-type"), "application/json")
}

// HasText checks if content type is text/plain
func (c *Client) HasText() bool {
	return c.Resp != nil && strings.Contains(c.Resp.Header.Get("content-type"), "text/plain")
}
