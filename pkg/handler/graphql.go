package handler

import (
	"fmt"
	"log"
	"net/http"
)

type GraphQLEndpoint struct {
	router Router
}

type Router interface {
	Post(query string) (string, error)
}

func NewGraphQLEndpoint(r Router) *GraphQLEndpoint {
	return &GraphQLEndpoint{
		router: r,
	}
}

func (e *GraphQLEndpoint) Handler(query string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		get, err := e.router.Post(query)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			_, err = w.Write([]byte(fmt.Errorf("querying the GraphQL endpoint: %w", err).Error()))
			if err != nil {
				log.Fatalf("writing the GraphQL result: %v", err)
			}
			return
		}
		_, err = w.Write([]byte(get))
		if err != nil {
			log.Fatalf("writing the GraphQL result: %v", err)
		}
	}
}
