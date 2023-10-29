package startgg

import (
	"context"
	"net/http"
	"strconv"

	"github.com/shurcooL/graphql"
)

type SGGClient struct {
	Client *graphql.Client
}

// CreateClient returns a SGGClient containing an authenticated graphql Client
func CreateClient(token string) SGGClient {
	httpClient := &http.Client{
		Transport: &authTransport{
			Token: token,
		},
	}
	c := graphql.NewClient("https://api.start.gg/gql/alpha", httpClient)

	return SGGClient{Client: c}
}

// =====================
// Queries
// =====================

// GetTournamentIdFromSlug returns the tournament Id given the friendly url string
func (c SGGClient) GetTournamentIdFromSlug(slug string) int {
	var query struct {
		Tournament struct {
			Id int
		} `graphql:"tournament(slug: $slug)"`
	}
	variables := map[string]any{
		"slug": graphql.String(slug),
	}

	err := c.Client.Query(context.Background(), &query, variables)

	if err != nil {
		panic(err)
	}

	return query.Tournament.Id
}

// GetTop8 returns a list of the Top 8 sets in a given event
func (c SGGClient) GetTop8(eventId int) []Node {
	var query struct {
		Event struct {
			Name string
			Sets struct {
				Nodes []Node
			} `graphql:"sets(page: $page, perPage: $perPage, sortType: STANDARD)"`
		} `graphql:"event(id: $eventId)"`
	}
	variables := map[string]any{
		"eventId": graphql.ID(strconv.Itoa(eventId)),
		"page":    graphql.Int(1),
		// 11 possible sets, including GF reset
		"perPage": graphql.Int(11),
	}

	err := c.Client.Query(context.Background(), &query, variables)

	if err != nil {
		panic(err)
	}

	sets := make([]Node, 0, 11)
	// only include sets where loser places top 8
	for _, node := range query.Event.Sets.Nodes {
		if node.LPlacement < 8 {
			sets = append(sets, node)
		}
	}

	return sets
}

// authTransport is a custom transport that adds the "Authorization" header to every request.
type authTransport struct {
	Token string
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Add the "Authorization" header to the request
	req.Header.Set("Authorization", "Bearer "+t.Token)

	// Use the default transport to execute the request
	return http.DefaultTransport.RoundTrip(req)
}
