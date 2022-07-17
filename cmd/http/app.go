package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pikomonde/i-view-nityo/cmd/initialize"
	d "github.com/pikomonde/i-view-nityo/delivery"
	rInmem "github.com/pikomonde/i-view-nityo/repository/inmem"
	s "github.com/pikomonde/i-view-nityo/service"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	logger := log.WithFields(log.Fields{})

	// Config
	config, err := initialize.NewConfig(ctx, logger)
	if err != nil {
		logger.WithField("err", err).Errorln("Failed to get config")
		return
	}

	// Initialization
	// init db

	// Repository
	repositoryUser := rInmem.NewRepositoryInMemUser(ctx, config)
	repositoryInvitation := rInmem.NewRepositoryInMemInvitation(ctx, config)

	// Service
	serviceLogin := s.NewServiceLogin(ctx, config, repositoryUser, repositoryInvitation)
	serviceInvitation := s.NewServiceInvitation(ctx, config, repositoryUser, repositoryInvitation)

	// Delivery
	delv := d.New(ctx, config, serviceLogin, serviceInvitation)
	if err := serviceLogin.RegisterAdminIfNotExist(); err != nil {
		panic(err)
	}
	delv.Start()

	term := make(chan os.Signal)
	signal.Notify(term, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-term:
		fmt.Println("Exiting gracefully...")
	}

}
