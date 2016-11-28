package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/kvstore/gen/client/kv"
)

// Default kvstore HTTP client.
var Default = NewHTTPClient(nil)

// NewHTTPClient creates a new kvstore HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Kvstore {
	if formats == nil {
		formats = strfmt.Default
	}
	transport := httptransport.New("localhost", "/", []string{"http"})
	return New(transport, formats)
}

// New creates a new kvstore client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Kvstore {
	cli := new(Kvstore)
	cli.Transport = transport

	cli.Kv = kv.New(transport, formats)

	return cli
}

// Kvstore is a client for kvstore
type Kvstore struct {
	Kv *kv.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Kvstore) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Kv.SetTransport(transport)

}
