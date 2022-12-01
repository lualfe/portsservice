package usecase

import (
	"context"
)

// Ports use case.
type Ports struct{}

// NewPorts creates a new Ports struct.
func NewPorts() *Ports {
	return &Ports{}
}

// Read reads the ports file and saves/updates its data to the database.
func (p *Ports) Read(ctx context.Context, file string) error {
	return nil
}
