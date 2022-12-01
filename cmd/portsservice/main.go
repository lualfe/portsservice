package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/lualfe/portsservice/internal/infrastructure/repository"
	"github.com/lualfe/portsservice/internal/usecase"
)

func main() {
	file := flag.String("filename", "", "the path and name of the file containing the ports")
	flag.Parse()

	db := repository.NewInMemory()

	portsUseCase := usecase.NewPorts(db)

	ctx := context.Background()

	if err := portsUseCase.Save(ctx, *file); err != nil {
		fmt.Println(err)
	}
}
