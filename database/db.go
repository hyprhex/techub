package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func DBConnection() {
	var (
		host = os.Getenv("HOST")
		port = os.Getenv("PORT")
		user = os.Getenv("USERNAME")
		pass = os.Getenv("PASSWORD")
		dbname = os.Getenv("DBNAME")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
	host, port, user, pass, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected")
}
