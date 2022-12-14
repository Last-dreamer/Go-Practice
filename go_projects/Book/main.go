package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"Book/controllers"
	"Book/driver"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

 
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

	db := driver.ConnectDB()
	controller := controllers.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/getBooks", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/getBook/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/updateBook", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/addNewBook", controller.AddBook(db)).Methods("POST")
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


 
 
 

