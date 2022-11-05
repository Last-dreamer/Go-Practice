package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)


var db *sql.DB

func ConnectDB() *sql.DB {
	pgUrl, err := pq.ParseURL(os.Getenv("SQL_URL"))
	logFetal(err)

	db, err = sql.Open("postgres",pgUrl)
	logFetal(err)
	
	err = db.Ping()
	logFetal(err)
	log.Println(pgUrl)

	return db
}


func logFetal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}