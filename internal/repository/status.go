// Package repository describes the interface of a Packet datasource
package repository

// Status exposes the interface to access a data source
type Status interface {
	CheckStatus() (*string, error)
}
