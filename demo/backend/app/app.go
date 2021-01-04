package app

import (
	"log"
	"net/http"
	"os"

	_handler "github.com/bluenumberfoundation/humanid-golang-demo/app/http"

	"github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	_humanID "github.com/bluenumberfoundation/humanid-golang-sdk"
)

// App is a
type App struct {
	Router  *mux.Router
}

// Initialize is a function to initializing the service
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
