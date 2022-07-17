package initialize_test

import (
	"context"
	"testing"

	"github.com/pikomonde/i-view-nityo/cmd/initialize"
	log "github.com/sirupsen/logrus"
)

func TestNewConfig(t *testing.T) {
	ctx := context.Background()
	logger := log.WithFields(log.Fields{})

	initialize.NewConfig(ctx, logger)
}
