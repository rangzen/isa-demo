package graphql

import (
	"io"
	"net/http"
	"strings"
)

type Repository struct {
	url   string
	query string
}

func NewRepository(url string, query string) *Repository {
	return &Repository{
		url:   url,
		query: query,
	}
}

func (r Repository) Query() (string, error) {
	resp, err := http.Post(
		r.url,
		"application/json",
		strings.NewReader(r.query),
	)
	if err != nil {
		return "", err
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
