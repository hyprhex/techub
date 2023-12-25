package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/hyprhex/techub/datatypes"
)

const jobUrl = "https://hacker-news.firebaseio.com/v0/item/%s.json"

func GetjobId() ([]int, error) {

	res, err := http.Get("https://hacker-news.firebaseio.com/v0/jobstories.json")
	if err != nil {
		return nil, err
	}

	var idList []int

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyBytes, &idList)

		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("Status code received: %v", res.StatusCode)
	}

	return idList, nil

}

func GetJobData(id int) (*datatypes.Job, error) {
	res, err := http.Get(fmt.Sprintf(jobUrl, strconv.Itoa(id)))
	if err != nil {
		return nil, err
	}

	var response JobResponse

	if res.StatusCode == http.StatusOK {
		bodybytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodybytes, &response)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("Status cide received: %v", res.StatusCode)
	}

	jobData := datatypes.Job{Title: response.Title, Url: response.URL}


	return &jobData, nil
}
