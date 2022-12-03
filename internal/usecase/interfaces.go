package usecase

import (
	"context"
	"io"

	"github.com/lualfe/portsservice/internal/entity"
	"github.com/lualfe/portsservice/pkg/portsstream"
)

// PortsRepo has operations on a repository of ports.
type PortsRepo interface {
	UpsertPort(ctx context.Context, port entity.Port) error
	BulkUpsertPort(ctx context.Context, ports []entity.Port) error
}

// PortsStreamer streams reader data.
type PortsStreamer interface {
	Start(ctx context.Context, r io.Reader)
	Stream() chan portsstream.Result
}
