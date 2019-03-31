package main

import (
	"github.com/go-mux-template/books"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}", books.CreateBook).Methods("POST")
	r.HandleFunc("/books/{title}", books.ReadBook).Methods("GET")
	r.HandleFunc("/books/{title}", books.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{title}", books.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8084", r))

}
