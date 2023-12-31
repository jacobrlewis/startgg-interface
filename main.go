package main

import (
	"fmt"
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
	fmt.Println(id)

	nodes := client.GetTop8(727876)
	fmt.Print(nodes)
}
