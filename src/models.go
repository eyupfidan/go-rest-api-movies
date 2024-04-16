package main

import (
	"math/rand"
	"time"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func init() {
	rand.Seed(time.Now().UnixNano())
	movies = []Movie{
		{ID: "1", Isbn: "456215", Title: "Avengers", Director: &Director{Firstname: "Stan", Lastname: "Lee"}},
		{ID: "2", Isbn: "435433", Title: "Green Mile", Director: &Director{Firstname: "Frank", Lastname: "Darabont"}},
	}
}
