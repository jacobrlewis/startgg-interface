package startgg

import (
	"context"
	"net/http"

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
		Tournament Tournament `graphql:"tournament(slug: $slug)"`
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
