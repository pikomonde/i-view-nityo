package httphandler

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/pikomonde/i-view-nityo/model"
	s "github.com/pikomonde/i-view-nityo/service"
	log "github.com/sirupsen/logrus"
)

// Handler is used to contains handler delivery
type Handler struct {
	Config            model.Config
	ServiceLogin      s.Login
	ServiceInvitation s.Invitation
	Mux               *http.ServeMux
}

// New returns new http handler
func New(
	config model.Config,
	serviceLogin s.Login,
	serviceInvitation s.Invitation,
) *Handler {
	return &Handler{
		ServiceLogin:      serviceLogin,
		ServiceInvitation: serviceInvitation,
		Config:            config,
		Mux:               http.NewServeMux(),
	}
}

// Start http server
func (h *Handler) Start() {
	h.RegisterPage()
	h.RegisterLogin()
	h.RegisterInvitation()

	// Starting server
	port := fmt.Sprintf(":%d", h.Config.App.Port)
	srv := &http.Server{
		ReadTimeout:  time.Duration(h.Config.App.ReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(h.Config.App.WriteTimeout) * time.Millisecond,
		TLSConfig: &tls.Config{
			PreferServerCipherSuites: true,
			CurvePreferences: []tls.CurveID{
				tls.CurveP256,
				tls.X25519,
			},
		},
		Handler: h.Mux,
		Addr:    port,
	}

	go func(srv *http.Server) {
		err := srv.ListenAndServe()
		if err != nil {
			log.Error(log.Fields{}, err)
		}
	}(srv)
	fmt.Println("Server listen to ", port)
}
