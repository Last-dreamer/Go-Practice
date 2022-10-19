package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"os"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
	// "strconv"
)

type Book struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Year int `json:"year"`
}

var allBooks  []Book
var db *sql.DB

func init(){
	gotenv.Load()
}


func logFetal(err error){
	if err != nil {
		log.Fatal(err)
	}
}
func main() {

	pgUrl, err := pq.ParseURL(os.Getenv("SQL_URL"))
	logFetal(err)

	db, err = sql.Open("postgres",pgUrl)
	logFetal(err)

	// log.Println(db)
	err = db.Ping()
	logFetal(err)

	log.Println(pgUrl)
	router := mux.NewRouter()

	router.HandleFunc("/getBooks", getBooks).Methods("GET")
	router.HandleFunc("/getBook/{id}", getBook).Methods("GET")
	router.HandleFunc("/updateBook", updateBook).Methods("PUT")
	router.HandleFunc("/addNewBook", addNewBook).Methods("POST")
	router.HandleFunc("/deleteBook/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", router))

}

func removeBook(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	result, _ := db.Exec("delete from  books where id=$1", params["id"])
	deleteResult, err := result.RowsAffected()
	logFetal(err)

	json.NewEncoder(w).Encode(deleteResult)

}


func updateBook(w http.ResponseWriter, r *http.Request){
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

  result, _ := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4", &book.Title, &book.Author, &book.Year, &book.ID)

  rowsUpdate, err := result.RowsAffected()
  logFetal(err)

  json.NewEncoder(w).Encode(rowsUpdate)

}

func getBook(w http.ResponseWriter, r *http.Request){
  var book Book
  params := mux.Vars(r)

  rows := db.QueryRow("select * from books where id=$1", params["id"])

  err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
  logFetal(err)

  json.NewEncoder(w).Encode(book)
}

func addNewBook(w http.ResponseWriter, r *http.Request) {
var book Book
var bookID int
json.NewDecoder(r.Body).Decode(&book)

err := db.QueryRow("insert into books(title, author, year) values($1, $2, $3) RETURNING id;",book.Title, book.Author, book.Year).Scan(&bookID)

logFetal(err)

json.NewEncoder(w).Encode(bookID)

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var book Book
	allBooks = []Book{}

	rows, err := db.Query("SELECT * FROM books")
	logFetal(err)

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFetal(err)

		allBooks = append(allBooks, book)
	}

	json.NewEncoder(w).Encode(allBooks)
}