package main

import (
	"log"
	"net/http"
)

func main() {
	r := NewRouter()
	log.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
