// Code generated by protoc-gen-goclay, but your can (must) modify it.
// source: api/strings.proto

package strings

import (
	desc "github.com/utrack/clay/integration/impl_type_name_tmpl/pb"
	"github.com/utrack/clay/v2/transport"
)

type Strings struct{}

func NewStrings() *Strings {
	return &Strings{}
}

// GetDescription is a simple alias to the ServiceDesc constructor.
// It makes it possible to register the service implementation @ the server.
func (i *Strings) GetDescription() transport.ServiceDesc {
	return desc.NewStringsServiceDesc(i)
}