package repository

import (
	"context"
)

// InMemory database
type InMemory map[string]any

// NewInMemory starts a new InMemory database.
func NewInMemory() InMemory {
	return make(InMemory)
}

// UpsertPorts inserts ports and updates them if they already exist.
func (i InMemory) UpsertPorts(ctx context.Context) error {
	return nil
}
