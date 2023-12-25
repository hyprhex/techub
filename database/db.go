package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)


var DB *sql.DB

func OpenDatabase() error {

	var (
		host = "localhost"
		port = 5432
		user = "postgres"
		pass = "root"
		dbname = "techub"
		err error
	)

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
	host, port, user, pass, dbname)
	DB, err = sql.Open("postgres", sqlInfo)
	if err != nil {
		return err
	}

	return nil
}

func CloseDatabse() error {
	return DB.Close()
}

func InsertRecord(title, url string) {
	sqlStatement := `
	INSERT INTO jobs (title, url)
	VALUES ($1, $2)
	`

	if err := DB.QueryRow(sqlStatement, title, url).Err(); err != nil {
		panic(err)
	}

	fmt.Println("Inserted successfully")
}
