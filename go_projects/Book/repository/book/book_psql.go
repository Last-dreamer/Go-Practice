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

func (b BookRepository) AddBook(db * sql.DB, book model.Book, bookID int) error {
	
	err := db.QueryRow("insert into books(title, author, year) values($1, $2, $3) RETURNING id;",book.Title, book.Author, book.Year).Scan(&bookID)

	if err != nil {
		return err
	}
	return nil
}

func (b BookRepository) UpdateBook(db *sql.DB, book model.Book) (int , error) {
	
	result, _ := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4", &book.Title, &book.Author, &book.Year, &book.ID)
	_, err := result.RowsAffected()
	if err != nil {
		return 0 , err
	}
	return 1, nil
}

func logFetal(err error) {
	panic("unimplemented")
}