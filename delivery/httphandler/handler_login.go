package httphandler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/pikomonde/i-view-nityo/model"
	s "github.com/pikomonde/i-view-nityo/service"
)

type LoginHandler struct {
	ServiceLogin s.Login
	Mux          *http.ServeMux
	Config       model.Config
	Clients      map[string]time.Time
}

func (h *Handler) RegisterLogin() {
	hh := LoginHandler{
		ServiceLogin: h.ServiceLogin,
		Mux:          h.Mux,
		Config:       h.Config,
		Clients:      make(map[string]time.Time),
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

func (hh *LoginHandler) rateLimiter(w http.ResponseWriter, r *http.Request) bool {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}

	clientID := fmt.Sprintf("%s::%s", ip, r.UserAgent())

	lastRequest, isExist := hh.Clients[clientID]
	if isExist {
		if lastRequest.Add(30 * time.Second).After(time.Now()) {
			return true
		}
	}

	hh.Clients[clientID] = time.Now()

	return false
}

func (hh *LoginHandler) LoginInvitation(w http.ResponseWriter, r *http.Request) {
	if isLimited := hh.rateLimiter(w, r); isLimited {
		respSuccessJSON(w, r, errorTryAgainLater)
		return
	}

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
