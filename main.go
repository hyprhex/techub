package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Job struct {
	By    string `json:"by"`
	ID    int    `json:"id"`
	Score int    `json:"score"`
	Text  string `json:"text"`
	Time  int    `json:"time"`
	Title string `json:"title"`
	Type  string `json:"type"`
	URL   string `json:"url"`
}

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

func main() {
	res, err := getJobsId(jobUrl)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range res {
        
        res, err := getJobData(jobItemUrl, v)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println(res)

	}

}
