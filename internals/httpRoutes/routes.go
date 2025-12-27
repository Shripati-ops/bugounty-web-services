package httpRoutes

import (
	"bugounty/internals/config"
	"bugounty/internals/controllers"
	"bugounty/internals/repositories"
	"bugounty/internals/services"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func SetupRoutes(config config.HTTPConfig, db *sql.DB) {
	r := mux.NewRouter()
	server := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%d", config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Server is running on port %d", config.Port)
	Routes(r, db)
	log.Fatal(server.ListenAndServe())
}

func Routes(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/", controllers.MainHandler).Methods("GET")

	// Declare your controllers, services, repos here separately.
	userRepo := repositories.NewUserRepository(db);
	userService := services.NewUserService(*userRepo)
	userHandler := controllers.NewUserHandler(userService)
	r.PathPrefix("/users").Handler(controllers.UserController(userHandler))
}
