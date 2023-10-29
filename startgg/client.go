package startgg

import (
	"context"
	"log"

	"github.com/machinebox/graphql"
)

type GQLClient struct {
	Client graphql.Client
	Bearer string
}

func (gql GQLClient) GetSets(eventId int, page int, perPage int) Event {

	query := EventSets

	request := graphql.NewRequest(query)
	request.Header.Add("Authorization", gql.Bearer)
	request.Var("eventId", eventId)
	request.Var("page", page)
	request.Var("perPage", perPage)

	var response struct {
		Event Event
	}
	err := gql.Client.Run(context.Background(), request, &response)
	if err != nil {
		panic(err)
	}

	return response.Event
}

func (gql GQLClient) GetEventName(eventId int) Event {
	query := EventName

	request := graphql.NewRequest(query)
	request.Header.Add("Authorization", gql.Bearer)
	request.Var("eventId", eventId)

	var response struct {
		Event Event
	}
	err := gql.Client.Run(context.Background(), request, &response)
	if err != nil {
		panic(err)
	}

	log.Print(response)
	return response.Event
}

func (gql GQLClient) GetTournamentIdFromSlug(slug string) int {
	query := TournamentIdFromSlug

	request := graphql.NewRequest(query)
	request.Header.Add("Authorization", gql.Bearer)
	request.Var("slug", slug)

	var response struct {
		Tournament Tournament
	}

	err := gql.Client.Run(context.Background(), request, &response)
	if err != nil {
		panic(err)
	}

	return response.Tournament.Id
}

func (gql GQLClient) GetTop8Contestants() {
	// TODO
}
