// Code generated by protoc-gen-goclay, but you can (must) modify it.
// source: strings.proto

package strings

import (
	desc "github.com/utrack/clay/integration/go_package_separate_impl_exists/pkg/strings"
	"github.com/utrack/clay/v3/transport"
)

type StringsImplementation struct {
	desc.UnimplementedStringsServer
}

// NewStrings create new StringsImplementation
func NewStrings() *StringsImplementation {
	return &StringsImplementation{}
}

// GetDescription is a simple alias to the ServiceDesc constructor.
// It makes it possible to register the service implementation @ the server.
func (i *StringsImplementation) GetDescription() transport.ServiceDesc {
	return desc.NewStringsServiceDesc(i)
}
