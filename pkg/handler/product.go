package handler

import (
	"log"
	"net/http"
)

type Product struct {
	rep ProductRepository
}

type ProductRepository interface {
	Query() (string, error)
}

func NewProduct(repository ProductRepository) *Product {
	return &Product{
		rep: repository,
	}
}

func (t *Product) Query() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		get, err := t.rep.Query()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				log.Fatalf("writing the Product Query page: %v", err)
			}
			return
		}
		_, err = w.Write([]byte(get))
		if err != nil {
			log.Fatalf("writing the Product Query page: %v", err)
		}
	}
}
