package main

import (
	"fmt"
	"log"
	"os"

	"github.com/IshanSaha05/IndiaVotes/pkg/mongodb"
	"github.com/IshanSaha05/IndiaVotes/pkg/scraper"
)

func main() {
	fmt.Println("Message: Connecting to MongoDb Server.")

	var mongoDBObject mongodb.MongoDB
	err := mongoDBObject.GetMongoClient()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	scraper.ScrapeStoreACData(&mongoDBObject)

	fmt.Println("Message: Job Done Successfully.")
}
