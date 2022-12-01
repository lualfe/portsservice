package repository

import (
	"context"

	"github.com/lualfe/portsservice/internal/entity"
)

// InMemory database
type InMemory map[string]entity.Port

// NewInMemory starts a new InMemory database.
func NewInMemory() InMemory {
	return make(InMemory)
}

// UpsertPort inserts a port and updates it if it already exists.
func (i InMemory) UpsertPort(ctx context.Context, port entity.Port) error {
	i[port.Key] = port
	return nil
}
