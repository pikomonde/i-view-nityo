package httphandler

import (
	"net/http"

	"github.com/pikomonde/i-view-nityo/model"
	s "github.com/pikomonde/i-view-nityo/service"
)

type InvitationHandler struct {
	ServiceInvitation s.Invitation
	Mux               *http.ServeMux
	Config            model.Config
}

func (h *Handler) RegisterInvitation() {
	hh := InvitationHandler{
		ServiceInvitation: h.ServiceInvitation,
		Mux:               h.Mux,
		Config:            h.Config,
	}

	hh.Mux.HandleFunc("/api/invite", hh.Invitation)
}

func (hh *InvitationHandler) Invitation(w http.ResponseWriter, r *http.Request) {
	err := hh.ServiceInvitation.CreateInvitation(10)
	if err != nil {
		respSuccessJSON(w, r, err.Error())
		return
	}
	respSuccessJSON(w, r, "success")
}
