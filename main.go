package main

import (
	"jacobrlewis/startgg-interface/startgg"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/machinebox/graphql"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	client := graphql.NewClient("https://api.start.gg/gql/alpha")
	gql := startgg.GQLClient{Client: *client, Bearer: "Bearer " + os.Getenv("api_key")}
	gql.GetSets(904060, 1, 1)
	gql.GetEventName(904060)
}
