package bookRepository

import (
	"Book/model"
	"database/sql"
)


type BookRepository struct {}

func (b BookRepository) GetBooks(db *sql.DB, book model.Book, allBooks []model.Book) ([]model.Book, error) {
  
	rows, err := db.Query("SELECT * FROM books")
	
	if err != nil {
		return []model.Book{}, err
	}
	
	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		allBooks = append(allBooks, book)
	}
	// json.NewEncoder(w).Encode(allBooks)
	if err != nil {
		return []model.Book{}, err
	}
	return allBooks, nil

}

func (b BookRepository) GetBook(db *sql.DB, book model.Book, id int) (model.Book, error) {
	rows := db.QueryRow("select * from books where id=$1", id)  
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	return book, err
}

func logFetal(err error) {
	panic("unimplemented")
}