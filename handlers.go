package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for _, movie := range movies {
		if movie.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	http.NotFound(w, r)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var updatedMovie Movie
	if err := json.NewDecoder(r.Body).Decode(&updatedMovie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedMovie.ID = id
	for i, movie := range movies {
		if movie.ID == id {
			movies[i] = updatedMovie
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedMovie)
			return
		}
	}
	http.NotFound(w, r)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for i, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:i], movies[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
	http.NotFound(w, r)
}
