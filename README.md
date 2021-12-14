
# humanID SDK for Go
humanid-golang-sdk is the official sdk for [humanID](https://human-id.org/).

Content:
- [Getting Started](#getting-started)
- [Examples](#sample-codes-using-"net/http"-library)
- [Contributing](#contributing)


## Getting Started

### Installation
```
go get github.com/bluenumberfoundation/humanid-golang-sdk
```
For the latest version of the SDK:
```
go get -u github.com/bluenumberfoundation/humanid-golang-sdk
```
### Dependencies
- "encoding/json"
- "io"
- "net/http"
- "bytes"

## Sample Codes using "net/http" library
#### main.go
```
package main

import (
	app "[PATH_TO_APP.GO]"
)

func main() {
	app := &app.App{}
	app.Initialize()

	app.Run(":[PORT_NUMBER]")
}

```
#### app.go

```
package app

import (
	"net/http"
	"os"

	_handler "[PATH_TO_YOUR_HTTP_HANDLER]"

	"github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	_humanID "github.com/bluenumberfoundation/humanid-golang-sdk"
)

type App struct {
	Router  *mux.Router
}

func (app *App) Initialize() {
	app.Router = mux.NewRouter()

	humanID := _humanID.New(
    os.Getenv("SERVER_ID"),
    os.Getenv("SERVER_SECRET"),
  )

	_handler.NewHandler(app.Router, humanID)
}


func (app *App) Run(addr string) {
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedHeaders([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
}

```
#### handler.go
```
package http

import (
	"encoding/json"
	"net/http"
	"fmt"

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

func (handler HTTPHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	tokenArr, ok := r.URL.Query()["et"]
	if !ok {
		code, ok := r.URL.Query()["code"]
		if !ok {
			redirectSuccessURL := "[YOUR_FAIL_REDIRECT_URL]"
			http.Redirect(w, r, redirectSuccessURL, 301)
			return
		}
	}
	token := tokenArr[0]

	verifyTokenResp, err := handler.humanID.VerifyToken(token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if verifyTokenResp.Code == "OK" {
		redirectSuccessURL := "[YOUR_SUCCCESS_REDIRECT_URL]"
		http.Redirect(w, r, redirectSuccessURL, 301)
		return
	} else {
		redirectSuccessURL := "[YOUR_FAIL_REDIRECT_URL]"
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
}

```
#### Environment Variables
These variables can be found in the Human ID Developer Console
```
SERVER_ID=[REPLACE_ME]
SERVER_SECRET=[REPLACE_ME]
CLIENT_ID=[REPLACE_ME]
CLIENT_SECRET=[REPLACE_ME]
```

You can access the full demo source file in the demo folder
## Contributing
### Run Tests
#### Pre-requisites (only for testing):
- docker
- docker-compose

#### Run tests:
```
docker-compose up
```
## License
Copyright 2019-2020 Bluenumber Foundation\
Licensed under the GNU General Public License v3.0 [(LICENSE)](https://github.com/human-internet/humanid-golang-sdk/blob/master/LICENSE)
