package httphandler

import (
	"net/http"

	s "github.com/pikomonde/i-view-nityo/service"
)

type LoginHandler struct {
	ServiceLogin s.Login
	Mux          *http.ServeMux
}

func (h *Handler) RegisterLogin() {
	hh := LoginHandler{
		ServiceLogin: h.ServiceLogin,
		Mux:          h.Mux,
	}

	hh.Mux.HandleFunc("/login", hh.Login)
}

func (hh *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	err := hh.ServiceLogin.LoginByInvitationID("")
	if err != nil {
		respSuccessJSON(w, r, err.Error())
		return
	}
	respSuccessJSON(w, r, "success")
}
