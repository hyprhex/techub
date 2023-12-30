package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	// "time"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Job struct {
	By    string `json:"by"`
	ID    int    `json:"id"`
	Score int    `json:"score"`
	Text  string `json:"text"`
	Time  int64  `json:"time"`
	Title string `json:"title"`
	Type  string `json:"type"`
	URL   string `json:"url"`
}

const (
	jobUrl     = "https://hacker-news.firebaseio.com/v0/jobstories.json"
	jobItemUrl = "https://hacker-news.firebaseio.com/v0/item/%s.json"
	jobUrlById = "https://news.ycombinator.com/item?id=%s"
)

func getJobsId(url string) ([]int, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var resIds []int

	if res.StatusCode == http.StatusOK {
		resData, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(resData, &resIds)

		if err != nil {
			return nil, err
		}

	}

	return resIds, nil
}

func getJobData(url string, id int) (*Job, error) {
	res, err := http.Get(fmt.Sprintf(url, strconv.Itoa(id)))
	if err != nil {
		return nil, err
	}

	var jobData Job

	if res.StatusCode == http.StatusOK {
		resData, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(resData, &jobData)
		if err != nil {
			return nil, err
		}
	}

	return &jobData, nil

}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func conectToDB() (*sql.DB, error) {
	var (
		host = os.Getenv("DB_HOST")
		port = os.Getenv("DB_PORT")
		user = os.Getenv("DB_USER")
		pass = os.Getenv("DB_PASS")
		name = os.Getenv("DB_NAME")
	)

	connectionString := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name)

	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil

}

func main() {
	db, err := conectToDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	res, err := getJobsId(jobUrl)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range res {

		res, err := getJobData(jobItemUrl, v)
		if err != nil {
			log.Fatal(err)
		}

		// newTime := time.Unix(res.Time, 0)

		if res.URL == "" {
			res.URL = fmt.Sprintf(jobUrlById, strconv.Itoa(res.ID))
		}

		getId := `
		SELECT jobid from jobs where jobid = $1
		`
		jobid := 0
		err = db.QueryRow(getId, res.ID).Scan(&jobid)

		if res.ID != jobid {
			sqlStatement := `
				INSERT INTO jobs (jobid, title, url, time)
				VALUES ($1, $2, $3, $4)
				RETURNING id`
			id := 0
			err = db.QueryRow(sqlStatement, res.ID, res.Title, res.URL, res.Time).Scan(&id)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(id)
		}

	}

}
