package portsstream

import (
	"github.com/lualfe/portsservice/internal/entity"
)

// Stream .
type Stream struct {
	stream chan Result
}

// Result ~.
type Result struct {
	Err  error
	Data entity.Port
}

// NewPortsStream creates a new Stream.
func NewPortsStream() *Stream {
	return &Stream{
		stream: make(chan Result),
	}
}
