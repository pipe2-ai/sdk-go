package pipe2

import (
	"net/http"

	"github.com/Khan/genqlient/graphql"
)

// NewClient creates a Pipe2.ai GraphQL client authenticated with the given JWT token.
func NewClient(token string, endpoint ...string) graphql.Client {
	url := "https://api.pipe2.ai/v1/graphql"
	if len(endpoint) > 0 && endpoint[0] != "" {
		url = endpoint[0]
	}
	return graphql.NewClient(url, &http.Client{
		Transport: &authTransport{token: token},
	})
}

type authTransport struct {
	token string
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+t.token)
	req.Header.Set("Content-Type", "application/json")
	return http.DefaultTransport.RoundTrip(req)
}
