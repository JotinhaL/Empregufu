package routes

import (
	"empregufu/pkg/controllers"
	"github.com/gorilla/mux"
)

var JobRoutes = func (router *mux.Router) {
	router.HandleFunc("/api/v1/jobs", controllers.CreateJob).Methods("POST")
	router.HandleFunc("/api/v1/jobs", controllers.GetAllJobs).Methods("GET")
	router.HandleFunc("/api/v1/jobs/{id}", controllers.GetJob).Methods("GET")
	router.HandleFunc("/api/v1/jobs/{id}", controllers.UpdateJob).Methods("PUT")
	router.HandleFunc("/api/v1/jobs/{id}", controllers.DeleteJob).Methods("DELETE")
}