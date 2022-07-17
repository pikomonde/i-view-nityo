package httphandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"text/template"

	"github.com/dgrijalva/jwt-go"
	"github.com/pikomonde/i-view-nityo/model"
	log "github.com/sirupsen/logrus"
)

type responseAPI struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func respSuccessJSON(w http.ResponseWriter, r *http.Request, data interface{}) {
	js, err := json.Marshal(responseAPI{
		Status: http.StatusOK,
		Data:   data,
	})
	if err != nil {
		respErrorText(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func respErrorJSON(w http.ResponseWriter, r *http.Request, status int, errStr string) {
	js, err := json.Marshal(responseAPI{
		Status: status,
		Data: struct {
			Message string `json:"message"`
		}{errStr},
	})
	if err != nil {
		respErrorText(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}

func respErrorText(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Something bad happened!"))
}

func respHTML(w http.ResponseWriter, r *http.Request, filename string, data map[string]interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	htmlFilepath := path.Join(defaultFilepath, filename)

	var tmpl, err = template.ParseFiles(htmlFilepath)
	if err != nil {
		log.Errorln("err: ", err)
		respErrorText(w, r)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Errorln("err: ", err)
		respErrorText(w, r)
	}
}

func parseJWT(w http.ResponseWriter, r *http.Request, jwtSecret string) (model.User, int, string) {
	// reqToken := r.Header.Get("Authorization")
	// splitToken := strings.Split(reqToken, "Bearer ")
	// if len(splitToken) != 2 {
	// 	return model.User{}, http.StatusUnauthorized, errorDeformedJWTToken
	// }
	// reqToken = splitToken[1]

	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			return model.User{}, http.StatusUnauthorized, errorCookieNotFound
		}
		// For any other type of error, return a bad request status
		return model.User{}, http.StatusBadRequest, errorCookieNotFound
	}
	reqToken := c.Value

	token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		// validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return model.User{}, http.StatusUnauthorized, errorCredentialProblem
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return model.User{}, http.StatusUnauthorized, errorExpiredJWTToken
	}

	user := model.User{}
	userStr, ok := claims["data"].(string)
	if !ok {
		return model.User{}, http.StatusUnauthorized, errorMissingJWTData
	}

	err = json.Unmarshal([]byte(userStr), &user)
	if err != nil {
		return model.User{}, http.StatusUnauthorized, errorDeformedJWTToken
	}

	return user, http.StatusOK, ""
}

func parseInput(w http.ResponseWriter, r *http.Request, v interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErrorJSON(w, r, http.StatusBadRequest, errBadRequest)
		return
	}

	err = json.Unmarshal(body, &v)
	if err != nil {
		respErrorJSON(w, r, http.StatusBadRequest, errBadRequest)
		return
	}
}

const (
	defaultFilepath = "delivery/httphandler/page/"

	errBadRequest          = "Bad request"
	errorUnauthorized      = "Unauthorized"
	errorCookieNotFound    = "Cookie Not Found"
	errorCredentialProblem = "Credential Problem"
	errorExpiredJWTToken   = "Expired JWT Token"
	errorMissingJWTData    = "Missing JWT Data"
	errorDeformedJWTToken  = "Deformed JWT Token"
)
