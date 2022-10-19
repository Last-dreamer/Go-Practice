package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Year string `json:"year"`
}


var allBooks []Book

func main(){

	allBooks = append(allBooks, 
		Book{ID: 1, Title: "Dream year", Author: "Dreamer", Year: "1999"},
		Book{ID: 2, Title: "Dream year2", Author: "Dreamer", Year: "2010"},
		Book{ID: 3, Title: "Dream year3", Author: "Dreamer", Year: "2021"},
	 )
	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}



func removeBook(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	for i, book := range allBooks {
		if book.ID == id {
			allBooks = append(allBooks[:i], allBooks[i+1:]...)
		}
	}

}

func updateBook(w http.ResponseWriter, r *http.Request){
	var book Book

	json.NewDecoder(r.Body).Decode(&book)

	for i, item := range allBooks {
		if item.ID == book.ID {
			allBooks[i] = book
		}
	}

	json.NewEncoder(w).Encode(allBooks)
}

func addBook(w http.ResponseWriter, r *http.Request){
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	allBooks = append(allBooks, book)
	json.NewEncoder(w).Encode(allBooks)
}

func getBook(w http.ResponseWriter, read *http.Request){
	param := mux.Vars(read)
 
	id, err := strconv.Atoi(param["id"])

	for _, book := range allBooks {
       
		if book.ID == id {
			json.NewEncoder(w).Encode(&book)
		} else {
			log.Println("error ",err)
		}
	}
}

func getBooks(w http.ResponseWriter, r *http.Request){
	 json.NewEncoder(w).Encode(allBooks)
  }