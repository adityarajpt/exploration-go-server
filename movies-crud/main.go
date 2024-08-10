package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"Director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatedMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&updatedMovie)

	params := mux.Vars(r)

	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			updatedMovie.ID = params["id"]
			movies = append(movies, updatedMovie)
			break
		}
	}
	json.NewEncoder(w).Encode(updatedMovie)
}

func main() {
	router := mux.NewRouter()

	// Prefilling the movies slice
	movies = append(movies, Movie{ID: "1", Isbn: "34212", Title: "First Book", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "56423", Title: "Second Book", Director: &Director{Firstname: "Jenny", Lastname: "Dummy"}})

	router.HandleFunc("/movies", getMovies).Methods(http.MethodGet)
	router.HandleFunc("/movies/{id}", getMovie).Methods(http.MethodGet)
	router.HandleFunc("movies", createMovie).Methods(http.MethodPost)
	router.HandleFunc("/movies/{id}", deleteMovie).Methods(http.MethodDelete)
	router.HandleFunc("/movies/{id}", updateMovie).Methods(http.MethodPut)

	fmt.Printf("Starting server at port:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
