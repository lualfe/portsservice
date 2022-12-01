package usecase

import (
	"context"
	"os"

	"github.com/lualfe/portsservice/pkg/portsstream"
)

// Ports use case.
type Ports struct {
	streamer PortsStreamer
	repo     PortsRepo
}

// NewPorts creates a new Ports use case.
func NewPorts(repo PortsRepo) *Ports {
	st := portsstream.NewPortsStream()
	return &Ports{
		streamer: st,
		repo:     repo,
	}
}

// Save reads ports from a file stream and saves it in the repository.
func (p *Ports) Save(ctx context.Context, file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	defer f.Close()

	errChan := make(chan error)
	go func() {
		defer close(errChan)
		for result := range p.streamer.Stream() {
			if result.Err != nil {
				errChan <- result.Err
				return
			}
			if err = p.repo.UpsertPort(ctx, result.Data); err != nil {
				errChan <- err
				return
			}
		}
	}()

	p.streamer.Start(ctx, f)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}
