package httphandler

import (
	"net/http"

	s "github.com/pikomonde/i-view-nityo/service"
)

type InvitationHandler struct {
	ServiceInvitation s.Invitation
	Mux               *http.ServeMux
}

func (h *Handler) RegisterInvitation() {
	hh := InvitationHandler{
		ServiceInvitation: h.ServiceInvitation,
		Mux:               h.Mux,
	}

	hh.Mux.HandleFunc("/invite", hh.Invitation)
}

func (hh *InvitationHandler) Invitation(w http.ResponseWriter, r *http.Request) {
	err := hh.ServiceInvitation.CreateInvitation(10)
	if err != nil {
		respSuccessJSON(w, r, err.Error())
		return
	}
	respSuccessJSON(w, r, "success")
}
