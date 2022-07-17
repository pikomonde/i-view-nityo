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

	hh.Mux.HandleFunc("/api/invitation/create", hh.CreateInvitation)
	hh.Mux.HandleFunc("/api/invitation/list", hh.ListInvitation)
	hh.Mux.HandleFunc("/api/invitation/disable", hh.DisableInvitation)
}

func (hh *InvitationHandler) CreateInvitation(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		respErrorJSON(w, r, http.StatusBadRequest, errBadRequest)
		return
	}

	user, status, errStr := parseJWT(w, r, hh.Config.App.JWTSecret)
	if status != http.StatusOK {
		respErrorJSON(w, r, status, errStr)
		return
	}
	if user.Role != model.UserRole_Admin {
		respErrorJSON(w, r, http.StatusUnauthorized, errorUnauthorized)
		return
	}

	invitation, err := hh.ServiceInvitation.CreateInvitation()
	if err != nil {
		respSuccessJSON(w, r, err.Error())
		return
	}
	respSuccessJSON(w, r, map[string]interface{}{
		"invitation_token": invitation.Token,
	})
}

func (hh *InvitationHandler) ListInvitation(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		respErrorJSON(w, r, http.StatusBadRequest, errBadRequest)
		return
	}

	user, status, errStr := parseJWT(w, r, hh.Config.App.JWTSecret)
	if status != http.StatusOK {
		respErrorJSON(w, r, status, errStr)
		return
	}
	if user.Role != model.UserRole_Admin {
		respErrorJSON(w, r, http.StatusUnauthorized, errorUnauthorized)
		return
	}

	invitations, err := hh.ServiceInvitation.GetInvitations()
	if err != nil {
		respSuccessJSON(w, r, err.Error())
		return
	}
	respSuccessJSON(w, r, map[string]interface{}{
		"invitations": invitations,
	})
}

func (hh *InvitationHandler) DisableInvitation(w http.ResponseWriter, r *http.Request) {
	var input struct {
		InvitationToken string `json:"invitation_token"`
	}

	if r.Method != "POST" {
		respErrorJSON(w, r, http.StatusBadRequest, errBadRequest)
		return
	}

	user, status, errStr := parseJWT(w, r, hh.Config.App.JWTSecret)
	if status != http.StatusOK {
		respErrorJSON(w, r, status, errStr)
		return
	}
	if user.Role != model.UserRole_Admin {
		respErrorJSON(w, r, http.StatusUnauthorized, errorUnauthorized)
		return
	}

	parseInput(w, r, &input)

	err := hh.ServiceInvitation.DisableInvitation(input.InvitationToken)
	if err != nil {
		respSuccessJSON(w, r, err.Error())
		return
	}
	respSuccessJSON(w, r, "success")
}
