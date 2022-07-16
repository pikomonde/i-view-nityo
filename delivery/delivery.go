package delivery

import (
	"context"

	"github.com/pikomonde/i-view-nityo/delivery/httphandler"
	"github.com/pikomonde/i-view-nityo/model"
	s "github.com/pikomonde/i-view-nityo/service"
)

type Delivery struct {
	HTTPHandler *httphandler.Handler
}

func New(
	ctx context.Context,
	config model.Config,
	serviceLogin s.Login,
	serviceInvitation s.Invitation,
) *Delivery {
	return &Delivery{
		HTTPHandler: httphandler.New(
			config,
			serviceLogin,
			serviceInvitation,
		),
	}
}

func (d *Delivery) Start() {
	d.HTTPHandler.Start()
}
