package graphql

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Router struct {
	url string
}

func NewRouter(url string) *Router {
	return &Router{
		url: url,
	}
}

func (r Router) Post(query string) (string, error) {
	resp, err := http.Post(
		r.url,
		"application/json",
		strings.NewReader(query),
	)
	if err != nil {
		return "", fmt.Errorf("posting the GraphQL query: %w", err)
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading the GraphQL response: %w", err)
	}
	return string(bytes), nil
}
