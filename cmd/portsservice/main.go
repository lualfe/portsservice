package main

import (
	"context"
	"fmt"

	"github.com/lualfe/portsservice/config"
	"github.com/lualfe/portsservice/internal/infrastructure/repository"
	"github.com/lualfe/portsservice/internal/usecase"
)

func main() {
	cfg := config.NewConfig()

	db := repository.NewInMemory(cfg.RedisAddress, cfg.RedisPassword)

	portsUseCase := usecase.NewPorts(db)

	ctx := context.Background()

	if err := portsUseCase.Save(ctx, cfg.FileName); err != nil {
		fmt.Println(err)
	}
}
