package controllers

import (
	"Book/model"
	bookRepository "Book/repository/book"
	"Book/utils"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Controller struct {} 

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book model.Book
		allBooks := []model.Book{}
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

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		var book model.Book
		var bookID int
		var error model.Error
		json.NewDecoder(r.Body).Decode(&book)
		bookRepo := bookRepository.BookRepository{}

		err := bookRepo.AddBook(db, book, bookID)
		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return 
		}

		w.Header().Set("Content-Type", "application/json")
		if bookID == 1 {
			utils.SendSuccess(w, "Successfully inserted ..")
		}
	
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		
		var book model.Book
		var bookRepo bookRepository.BookRepository
		var error model.Error
	    json.NewDecoder(r.Body).Decode(&book)

		if book.ID == 0 || book.Author == "" || book.Title == "" || book.Year == 0 {
			error.Message = "All Fields are required"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}
        

		isRowsUpdated, err := bookRepo.UpdateBook(db, book)
		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
		  return
		}

		w.Header().Set("Content-Type", "application/json")
		if isRowsUpdated == 1 {	
		  utils.SendSuccess(w, "Updated Successfully")   
		}
	}
}

func logFetal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
