package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v9"

	"github.com/lualfe/portsservice/internal/entity"
)

// InMemory database
type InMemory struct {
	client redisClient
}

type redisClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Pipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error)
}

// NewInMemory starts a new InMemory database.
func NewInMemory(address, password string) *InMemory {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
	})
	return &InMemory{
		client: rdb,
	}
}

// UpsertPort inserts a port and updates it if it already exists.
func (i *InMemory) UpsertPort(ctx context.Context, port entity.Port) error {
	data, _ := json.Marshal(port)
	cmd := i.client.Set(ctx, port.Key, string(data), 0)
	return cmd.Err()
}

func (i *InMemory) BulkUpsertPort(ctx context.Context, ports []entity.Port) error {
	_, err := i.client.Pipelined(ctx, func(pipeliner redis.Pipeliner) error {
		for _, p := range ports {
			data, _ := json.Marshal(p)
			pipeliner.Set(ctx, p.Key, string(data), 0)
		}
		return nil
	})

	return err
}
