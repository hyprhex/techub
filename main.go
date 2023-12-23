package main

import (
	"fmt"
	"log"

	"github.com/hyprhex/techub/api"
)

func main() {

	res, err := api.GetjobId()
	
	if err != nil {
		log.Fatal(err)
	}

	
	for _, n := range res {
		res, err := api.GetJobData(n)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(res.Title, res.Url)
	}


}
