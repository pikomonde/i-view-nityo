package httphandler

import (
	"net/http"

	"github.com/pikomonde/i-view-nityo/model"
)

type PageHandler struct {
	Mux    *http.ServeMux
	Config model.Config
}

func (h *Handler) RegisterPage() {
	hh := PageHandler{
		Mux:    h.Mux,
		Config: h.Config,
	}

	hh.Mux.HandleFunc("/", hh.Index)
	hh.Mux.HandleFunc("/admin", hh.Admin)

}

func (hh *PageHandler) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		respErrorJSON(w, r, http.StatusBadRequest, errBadRequest)
		return
	}

	user, status, _ := parseJWT(w, r, hh.Config.App.JWTSecret)
	if status == http.StatusOK {
		if user.Role == model.UserRole_Admin {
			respHTML(w, r, "dashboard.html", map[string]interface{}{})
			return
		} else if user.Role == model.UserRole_Invitation {
			respHTML(w, r, "invitation.html", map[string]interface{}{})
			return
		}
		respHTML(w, r, "unauthorized.html", map[string]interface{}{})
		return
	}
	respHTML(w, r, "index.html", map[string]interface{}{})
}

func (hh *PageHandler) Admin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		respErrorJSON(w, r, http.StatusBadRequest, errBadRequest)
		return
	}

	user, status, _ := parseJWT(w, r, hh.Config.App.JWTSecret)
	if status == http.StatusOK {
		if user.Role == model.UserRole_Admin {
			respHTML(w, r, "dashboard.html", map[string]interface{}{})
			return
		} else if user.Role == model.UserRole_Invitation {
			respHTML(w, r, "invitation.html", map[string]interface{}{})
			return
		}
		respHTML(w, r, "unauthorized.html", map[string]interface{}{})
		return
	}
	respHTML(w, r, "admin.html", map[string]interface{}{})
}
