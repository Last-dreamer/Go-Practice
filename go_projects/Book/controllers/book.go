package controllers

import (
	"Book/model"
	bookRepository "Book/repository/book"
	"Book/utils"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Controller struct {}


var allBooks  []model.Book

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book model.Book
		allBooks = []model.Book{}
		var error model.Error
		
	   getBooks := bookRepository.BookRepository{}
	   getAllBooks, err :=	getBooks.GetBooks(db, book, allBooks)

	   if err != nil {
		   error.Message  = "Server Error"
		   utils.SendError(w, http.StatusInternalServerError, error)
		   return
	   }

	   w.Header().Set("Content-Type", "application/json")
	   utils.SendSuccess(w,getAllBooks)
	}
}


func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	
	return func(w http.ResponseWriter, r *http.Request) {
		var book model.Book
		params := mux.Vars(r)
		var error model.Error
	    bookRepo := bookRepository.BookRepository{}

	   id, _ := strconv.Atoi(params["id"])

		getBook, err := bookRepo.GetBook(db, book, id)
		
		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "No Book Found"
			    utils.SendError(w, http.StatusNotFound, error)
				return
			}else{
				error.Message = "Server Error"
				utils.SendError(w, http.StatusInternalServerError, error)
				return 
			}
		}
		
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, getBook)
	 
	}
}

func logFetal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
