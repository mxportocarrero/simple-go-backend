package api

import (
	"net/http"
	v1 "simple-go-backend/internal/api/v1"
	"simple-go-backend/internal/database"

	"github.com/gorilla/mux"
)

func NewRouter(db database.Database) (http.Handler, error) {
	router := mux.NewRouter()

	router.HandleFunc("/version", v1.VersionHandler)

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	userAPI := &v1.UserAPI{
		DB: db,
	}

	apiRouter.HandleFunc("/users", userAPI.GetAllUsers).Methods("GET")
	apiRouter.HandleFunc("/users", userAPI.CreateUser).Methods("POST")
	apiRouter.HandleFunc("/users/{id}", userAPI.UpdateUser).Methods("PUT")

	return router, nil
}
