package pg

import (
	"database/sql"
	"encoding/json"
	. "github.com/go-jet/jet/v2/postgres"
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
	err := stmt.Query(r.db, &dest)
	if err != nil {
		return "", err
	}

	jsonTxt, err := json.MarshalIndent(dest, "", "  ")
	if err != nil {
		return "", err
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
	err := stmt.Query(r.db, &dest)
	if err != nil {
		return "", err
	}

	jsonTxt, err := json.MarshalIndent(dest, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonTxt), nil
}
