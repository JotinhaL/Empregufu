package main

import (
	"empregufu/pkg/models"
	"empregufu/pkg/routes"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	models.Init()
	r := mux.NewRouter()

	routes.JobRoutes(r)
	http.Handle("/", r)

	fmt.Println("Starting server on http://localhost:8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
  	   log.Fatalf("Failed to start server: %v", err)
 	   return
	} 

}
