package handler_test

import (
	"database/sql"
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

type CountryRepositoryNoRows struct{}

func (c CountryRepositoryNoRows) Get(_ string) (string, error) {
	return "", sql.ErrNoRows
}

func (c CountryRepositoryNoRows) GetAll() (string, error) {
	return "", sql.ErrNoRows
}

func TestCountry_Uni(t *testing.T) {

	t.Run("when repository error, then send back a 500", func(t *testing.T) {
		repError := h.NewCountry(CountryRepositoryError{})
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/country/england", nil)

		repError.Uni()(w, r)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("expected %d, got %d", http.StatusNotFound, w.Code)
		}
	})

	t.Run("when repository no rows, then send back a 404", func(t *testing.T) {
		repNoRows := h.NewCountry(CountryRepositoryNoRows{})
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/country/england", nil)

		repNoRows.Uni()(w, r)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected %d, got %d", http.StatusNotFound, w.Code)
		}
	})
}

func TestCountry_All(t *testing.T) {

	t.Run("when repository error, then send back a 500", func(t *testing.T) {
		repError := h.NewCountry(CountryRepositoryError{})
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/country", nil)

		repError.All()(w, r)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("expected %d, got %d", http.StatusNotFound, w.Code)
		}
	})

	t.Run("when repository no rows, then send back a 404", func(t *testing.T) {
		repNoRows := h.NewCountry(CountryRepositoryNoRows{})
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/country", nil)

		repNoRows.All()(w, r)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected %d, got %d", http.StatusNotFound, w.Code)
		}
	})
}
