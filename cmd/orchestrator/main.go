package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/DODIK22816/final5/internal/config"
	"github.com/DODIK22816/final5/internal/orchestrator/delivery/rest/servers"
	"github.com/DODIK22816/final5/internal/orchestrator/repository"
	"github.com/DODIK22816/final5/internal/orchestrator/service"
	"github.com/DODIK22816/final5/pkg/db/cache"
	"github.com/DODIK22816/final5/pkg/db/postgres"
	"github.com/DODIK22816/final5/pkg/utils/errors"
)

func main() {
	ctx := context.Background()
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	db, err := postgres.New(cfg.PostgresConfig)
	if err != nil {
		panic(err)
	}
	redis := cache.New(cfg.RedisConfig)
	fmt.Println(redis.Ping(ctx))

	repo := repository.NewCalculatorRepository(ctx, db, redis)
	srv := service.NewCalculatorService(repo)

	createAgentUser(repo)

	server := servers.NewCalculatorServer(cfg.CalculatorServerConfig, srv)

	go func() {
		if err := server.Start(); err != nil {
			log.Println(err)
		}
	}()

	gracefulShutdownChannel := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdownChannel, syscall.SIGTERM, syscall.SIGINT)

	<-gracefulShutdownChannel
	err = server.Stop(ctx)
	if err != nil {
		log.Println(err)
	}
}

func createAgentUser(repo *repository.CalculatorRepository) {
	err := repo.Register("agent", "agent_password")
	if err != nil && err != errors.ErrUserAlreadyExists {
		log.Printf("Error with creating an agent-user: %v", err)
	} else {
		log.Println("Agent-user has been successfully created or already exists")
	}
}
