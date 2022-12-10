package handler_test

import (
	"fmt"
	h "github.com/rangzen/isa-demo/pkg/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

type CountryRepositoryError struct{}

func (c CountryRepositoryError) Get(_ string) (string, error) {
	return "", fmt.Errorf("something went wrong")
}

func (c CountryRepositoryError) GetAll() (string, error) {
	return "", fmt.Errorf("something went wrong")
}

func TestCountry_Uni(t *testing.T) {
	errRep := h.NewCountry(CountryRepositoryError{})

	t.Run("when repository error, then send back a 404", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/country/england", nil)

		errRep.Uni()(w, r)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected %d, got %d", http.StatusNotFound, w.Code)
		}
	})
}

func TestCountry_All(t *testing.T) {
	errRep := h.NewCountry(CountryRepositoryError{})

	t.Run("when repository error, then send back a 404", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/country", nil)

		errRep.All()(w, r)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected %d, got %d", http.StatusNotFound, w.Code)
		}
	})
}
