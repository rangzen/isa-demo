package handler

import (
	"database/sql"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Country struct {
	rep CountryRepository
}

type CountryRepository interface {
	Get(country string) (string, error)
	GetAll() (string, error)
}

func NewCountry(repository CountryRepository) *Country {
	return &Country{
		rep: repository,
	}
}

func (c *Country) Uni() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		get, err := c.rep.Get(vars["country"])
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				w.WriteHeader(http.StatusNotFound)
			default:
				w.WriteHeader(http.StatusInternalServerError)
			}
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				log.Fatalf("writing the Country Uni page: %v", err)
			}
			return
		}
		_, err = w.Write([]byte(get))
		if err != nil {
			log.Fatalf("writing the Country Uni page: %v", err)
		}
	}
}

func (c *Country) All() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		getAll, err := c.rep.GetAll()
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				w.WriteHeader(http.StatusNotFound)
			default:
				w.WriteHeader(http.StatusInternalServerError)
			}
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				log.Fatalf("writing the Country All page: %v", err)
			}
			return
		}
		_, err = w.Write([]byte(getAll))
		if err != nil {
			log.Fatalf("writing the Country All page: %v", err)
		}
	}
}
