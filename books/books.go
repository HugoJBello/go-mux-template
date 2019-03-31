package books

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Book struct {
	page   string    `json:"page"`
	title   string    `json:"title"`
}


func CreateBook (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	Book := Book{title:title, page:"1"}
	fmt.Println(Book)
	respondWithJSON(w, 200, Book)
}


func ReadBook (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	Book := Book{title:title, page:"1"}
	fmt.Println(Book)

	respondWithJSON(w, 200, Book)
}

func UpdateBook (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	//fmt.Fprintf(w, "You've requested the book: %s", title)

	Book := Book{title:title, page:"1"}
	respondWithJSON(w, 200, Book)
}


func DeleteBook (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	Book := Book{title:title, page:"1"}
	respondWithJSON(w, 200, Book)
}



func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}