package main

import (
	"context"
	"log"

	"github.com/GoExpertCurso/auction-go/config/database/mongodb"
	"github.com/joho/godotenv"
)

func main() {

	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error trying to load .env file")
	}

	databaseClient, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
