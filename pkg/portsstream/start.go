package portsstream

import (
	"context"
	"encoding/json"
	"io"

	"github.com/lualfe/portsservice/internal/entity"
)

// Start reads the ports and sends them to the Result.
func (s *Stream) Start(ctx context.Context, r io.Reader) {
	defer close(s.stream)

	decoder := json.NewDecoder(r)

	var err error
	if _, err = decoder.Token(); err != nil {
		s.stream <- Result{Err: err}
		return
	}

	for decoder.More() {
		var port entity.Port

		portKey, err := decoder.Token()
		if err != nil {
			s.stream <- Result{Err: err}
			return
		}

		port.Key = portKey.(string)

		if err = decoder.Decode(&port); err != nil {
			s.stream <- Result{Err: err}
			return
		}

		s.stream <- Result{Data: port}
	}

	if _, err = decoder.Token(); err != nil {
		s.stream <- Result{Err: err}
		return
	}
}
