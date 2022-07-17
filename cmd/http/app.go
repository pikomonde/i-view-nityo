package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pikomonde/i-view-nityo/cmd/initialize"
	d "github.com/pikomonde/i-view-nityo/delivery"
	"github.com/pikomonde/i-view-nityo/model"

	rInmem "github.com/pikomonde/i-view-nityo/repository/inmem"
	rMySql "github.com/pikomonde/i-view-nityo/repository/mysql"
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
	mySQLCli, err := initialize.NewMySQL(config.MySQL)
	if err != nil {
		logger.WithField("err", err).Errorln("Failed to connect mysql")
		return
	}

	// Repository
	repositoryInMemUser := rInmem.NewRepositoryInMemUser(ctx, config)
	repositoryInMemInvitation := rInmem.NewRepositoryInMemInvitation(ctx, config)
	repositoryMySQLUser := rMySql.NewRepositoryInMemUser(ctx, config, mySQLCli)
	repositoryMySQLInvitation := rMySql.NewRepositoryInMemInvitation(ctx, config, mySQLCli)

	// Service
	serviceLogin := s.NewServiceLogin(ctx, config, repositoryInMemUser, repositoryInMemInvitation)
	serviceInvitation := s.NewServiceInvitation(ctx, config, repositoryInMemUser, repositoryInMemInvitation)
	if config.App.DBType == model.DBType_MySQL {
		serviceLogin = s.NewServiceLogin(ctx, config, repositoryMySQLUser, repositoryMySQLInvitation)
		serviceInvitation = s.NewServiceInvitation(ctx, config, repositoryMySQLUser, repositoryMySQLInvitation)
	}

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
