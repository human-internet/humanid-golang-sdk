package http

import (
	"encoding/json"
	"net/http"
	"fmt"
	// "strconv"
	// "os"

	"github.com/gorilla/mux"
	_humanID "github.com/bluenumberfoundation/humanid-golang-sdk"
)

type HTTPHandler struct {
	humanID _humanID.HumanID
}

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// https://human-id.org/?code=ERR_10&message=Invalid%20phone%20number%20input
// av9WfTOwZtnc6ibLaLffsSmA%2FJJ5av7pNZJergIE%2BtOW0SqBRdQhiuIakPMMphvv%2BgoX%2FQQ8gWnq9OKwZjYlQdts3fGEzrSCX48foKw7j5dO1QRhXNpdENrKS%2B8C3ziASiHzox8Wg3cJ162KJlzPjr7wz2P%2BCw3Akg40PymguJCTUWQ%3D%3D
func (handler HTTPHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	tokenArr, ok := r.URL.Query()["et"]
	if !ok {
		code, ok := r.URL.Query()["code"]
		if !ok {
			redirectSuccessURL := "http://www.example.com/fail"
			http.Redirect(w, r, redirectSuccessURL, 301)
			return
		}
	}
	token := tokenArr[0]

	fmt.Println(token)
	verifyTokenResp, err := handler.humanID.VerifyToken(token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if verifyTokenResp.Code == "OK" {
		redirectSuccessURL := "http://www.example.com/success"
		http.Redirect(w, r, redirectSuccessURL, 301)
		return
	} else {
		redirectSuccessURL := "http://www.example.com/fail"
		http.Redirect(w, r, redirectSuccessURL, 301)
		return
	}
}

func (handler HTTPHandler) RequestLoginURL(w http.ResponseWriter, r *http.Request) {
	loginResp, err := handler.humanID.Login("ID", "")
	if err != nil {
		payload := AuthenticateResponse{
			Success: false,
			Code: "500",
			Message: "Internal Server Error",
		}
		RespondJSON(w, http.StatusInternalServerError, payload)
		return
	}
	if loginResp.Code == "OK" {
		loginURL := loginResp.Data.WebLoginUrl
		http.Redirect(w, r, loginURL, 301)
		return
	}
	payload := AuthenticateResponse{
		Success: false,
		Code: "500",
		Message: "Internal Server Error",
	}
	RespondJSON(w, http.StatusInternalServerError, payload)
	return
}


func NewHandler(publicRouter *mux.Router, humanID _humanID.HumanID) {
	handler := &HTTPHandler{
		humanID: humanID,
	}

	publicRouter.HandleFunc("/authenticate", handler.Authenticate).Methods("GET")
	publicRouter.HandleFunc("/request", handler.RequestLoginURL).Methods("GET")
	// publicRouter.HandleFunc("/authenticate", handler.Authenticate).Methods("GET")
}
