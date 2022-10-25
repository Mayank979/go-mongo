package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mayankyadav/mongodb/helpers"
	model "github.com/mayankyadav/mongodb/models"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := helpers.GetAllItems()

	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie model.Netflix

	if r.Body == nil {
		log.Fatal("req body is empty")
	}

	_ = json.NewDecoder(r.Body).Decode(&movie)

	helpers.InsertOne(movie)

	json.NewEncoder(w).Encode(movie)

}

func MarcAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	helpers.UpdateOne(params["id"])

	json.NewEncoder(w).Encode(params)

}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	result := helpers.DeleteOne(params["id"])
	json.NewEncoder(w).Encode(result)
}

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result := helpers.DeleteAll()
	json.NewEncoder(w).Encode(result)
}
