package httpRoutes

import (
	"bugounty/internals/config"
	"bugounty/internals/controllers"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

func SetupRoutes(config config.HTTPConfig) {
	r := mux.NewRouter()
	server := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%d", config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Server is running on port %d", config.Port)
	Routes(r)
	log.Fatal(server.ListenAndServe())
}

func Routes(r *mux.Router) {
	r.HandleFunc("/", controllers.MainHandler).Methods("GET")
}
