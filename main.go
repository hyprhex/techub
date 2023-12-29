package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	jobUrl     = "https://hacker-news.firebaseio.com/v0/jobstories.json"
	jobItemUrl = "https://hacker-news.firebaseio.com/v0/item/%s.json"
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

func main() {
	res, err := getJobsId(jobUrl)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range res {
		fmt.Println(v)
	}

}
