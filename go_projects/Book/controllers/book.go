package controllers

import (
	"Book/model"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Controller struct {}


var allBooks  []model.Book

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book model.Book
		allBooks = []model.Book{}
	
		rows, err := db.Query("SELECT * FROM books")
		logFetal(err)
	
		for rows.Next() {
			err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
			logFetal(err)
			allBooks = append(allBooks, book)
		}
		json.NewEncoder(w).Encode(allBooks)
	}
}


func logFetal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
