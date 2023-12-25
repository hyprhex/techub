package main

import (
	"log"

	"github.com/hyprhex/techub/api"
	"github.com/hyprhex/techub/database"
)

func main() {
	
	if err := database.OpenDatabase(); err != nil {
		log.Printf("Error connecting to DB %v", err)
	}
	
	defer database.CloseDatabse()

	res, err := api.GetjobId()

	if err != nil {
		log.Fatal(err)
	}


	for _, n := range res {
		res, err := api.GetJobData(n)
		if err != nil {
			log.Fatal(err)
		}
	
		database.InsertRecord(res.Title, res.Url)
	}

}
