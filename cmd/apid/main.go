package main

import (
	"log"

	"github.com/jbelmont/api-workshop/cmd/apid/config"
)

func main() {
	config := config.New()

	log.Println(config)

	log.Println("main : Started : Initialize Mongo")
	// masterDB, err := mongo.New(config.MongoHost, config.DBTimeout)
	// if err != nil {
	// 	log.Fatalf("startup : Register DB : %v", err)
	// }
	// defer masterDB.Close()

	// server := http.Server{
	// 	Addr:           config.APIHost,
	// 	Handler:        handlers.API(masterDB, config),
	// 	ReadTimeout:    config.ReadTimeout,
	// 	WriteTimeout:   config.WriteTimeout,
	// 	MaxHeaderBytes: 1 << 20,
	// }
}
