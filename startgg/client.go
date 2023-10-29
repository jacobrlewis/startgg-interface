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

func (gql GQLClient) GetSets(eventId int, page int, perPage int) EventResponse {

	query := EventSets

	request := graphql.NewRequest(query)
	request.Header.Add("Authorization", gql.Bearer)
	request.Var("eventId", eventId)
	request.Var("page", page)
	request.Var("perPage", perPage)

	var response EventResponse
	err := gql.Client.Run(context.Background(), request, &response)
	if err != nil {
		panic(err)
	}

	log.Print(response)
	log.Print(response.Event.Id)
	log.Print(response.Event.Name)

	for _, node := range response.Event.Sets.Nodes {
		for _, slot := range node.Slots {
			log.Print(slot.Entrant.Name)
		}
	}
	return response
}

func (gql GQLClient) GetEventName(eventId int) EventResponse{
	query := EventName

	request := graphql.NewRequest(query)
	request.Header.Add("Authorization", gql.Bearer)
	request.Var("eventId", eventId)

	var response EventResponse
	err := gql.Client.Run(context.Background(), request, &response)
	if err != nil {
		panic(err)
	}

	log.Print(response)
	return response
}

func (gql GQLClient) GetTop8Contestants() {
	// TODO
}