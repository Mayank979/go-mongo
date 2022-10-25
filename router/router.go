package router

import (
	"fmt"

	"github.com/gorilla/mux"
	controller "github.com/mayankyadav/mongodb/controllers"
)

func Router() *mux.Router {
	fmt.Println("inside router")
	r := mux.NewRouter()

	r.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movie/{id}", controller.CreateMovie).Methods("PUT")
	r.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	r.HandleFunc("/api/delete", controller.DeleteAll).Methods("DELETE")

	return r
}
