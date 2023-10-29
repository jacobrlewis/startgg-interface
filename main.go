package main

import (
	"jacobrlewis/startgg-interface/startgg"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	client := startgg.CreateClient(os.Getenv("api_key"))

	id := client.GetTournamentIdFromSlug("genesis-x")
	log.Println(id)
}
