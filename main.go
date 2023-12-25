package main

import (
	"fmt"
	"log"

	"github.com/hyprhex/techub/api"
	"github.com/hyprhex/techub/database"
)

func main() {

	database.DBConnection()

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
