package httphandler

import (
	"net/http"

	"github.com/pikomonde/i-view-nityo/model"
	s "github.com/pikomonde/i-view-nityo/service"
)

type LoginHandler struct {
	ServiceLogin s.Login
	Mux          *http.ServeMux
	Config       model.Config
}

func (h *Handler) RegisterLogin() {
	hh := LoginHandler{
		ServiceLogin: h.ServiceLogin,
		Mux:          h.Mux,
		Config:       h.Config,
	}

	hh.Mux.HandleFunc("/api/login", hh.Login)
	hh.Mux.HandleFunc("/api/login-invitation", hh.LoginInvitation)
}

func (hh *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if r.Method != "POST" {
		respErrorJSON(w, r, http.StatusBadRequest, errBadRequest)
		return
	}

	parseInput(w, r, &input)

	token, err := hh.ServiceLogin.LoginByUsernamePassword(input.Username, input.Password)
	if err != nil {
		respSuccessJSON(w, r, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: token,
	})
	respSuccessJSON(w, r, map[string]interface{}{
		"token": token,
	})
}

func (hh *LoginHandler) LoginInvitation(w http.ResponseWriter, r *http.Request) {
	var input struct {
		InvitationToken string `json:"invitation_token"`
	}

	if r.Method != "POST" {
		respErrorJSON(w, r, http.StatusBadRequest, errBadRequest)
		return
	}

	parseInput(w, r, &input)

	token, err := hh.ServiceLogin.LoginByInvitationToken(input.InvitationToken)
	if err != nil {
		respSuccessJSON(w, r, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: token,
	})
	respSuccessJSON(w, r, map[string]interface{}{
		"token": token,
	})
}
