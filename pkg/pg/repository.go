package pg

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	. "github.com/rangzen/isa-demo/.gen/postgres/public/table"
)

type Answer struct {
	Country string
	Capital string
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r Repository) Get(country string) (string, error) {
	stmt := SELECT(
		Country.Name.AS("answer.country"),
		City.Name.AS("answer.capital"),
	).FROM(
		Country.
			INNER_JOIN(City, Country.Capital.EQ(City.ID)),
	).WHERE(
		Country.Name.EQ(String(country)),
	)

	var dest Answer

	if err := stmt.Query(r.db, &dest); err != nil {
		switch {
		case errors.Is(err, qrm.ErrNoRows):
			return "", sql.ErrNoRows
		default:
			return "", fmt.Errorf("querying the database: %w", err)
		}
	}

	jsonTxt, err := json.MarshalIndent(dest, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshalling the JSON: %w", err)
	}

	return string(jsonTxt), nil
}

func (r Repository) GetAll() (string, error) {
	stmt := SELECT(
		Country.Name.AS("answer.country"),
		City.Name.AS("answer.capital"),
	).FROM(
		Country.
			INNER_JOIN(City, Country.Capital.EQ(City.ID)),
	).ORDER_BY(
		Country.Name,
	)

	var dest []Answer

	if err := stmt.Query(r.db, &dest); err != nil {
		switch {
		case errors.Is(err, qrm.ErrNoRows):
			return "", sql.ErrNoRows
		default:
			return "", fmt.Errorf("querying the database: %w", err)
		}
	}

	jsonTxt, err := json.MarshalIndent(dest, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshalling the JSON: %w", err)
	}

	return string(jsonTxt), nil
}
