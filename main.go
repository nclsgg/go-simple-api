package main

import (
	"FirstAPI/cmd/api"
	"FirstAPI/internal/config"
	"FirstAPI/internal/infra/db"
)

func init() {
	config.LoadEnv()
}

func main() {
	dbConnector, mongoErr := db.NewMongoDBService()
	if mongoErr != nil {
		panic(mongoErr)
	}

	api.NewApi(dbConnector).Start()
}
