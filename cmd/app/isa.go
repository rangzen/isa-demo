package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rangzen/isa-demo/pkg/handler"
	"github.com/rangzen/isa-demo/pkg/pg"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	loadEnvVar()

	// PostgreSQL
	pgConn := generatePgConnData()
	pgDb, err := sql.Open("postgres", pgConn)
	if err != nil {
		log.Fatalf("error opening PostgreSQL connection: %v", err)
	}
	defer func(pgDb *sql.DB) {
		err := pgDb.Close()
		if err != nil {
			log.Fatalf("error closing the PostgreSQL database: %v", err)
		}
	}(pgDb)

	pgRepo := pg.NewRepository(pgDb)
	pgCountryHandler := handler.NewCountry(pgRepo)

	// Router
	r := mux.NewRouter()
	r.HandleFunc("/", handler.Home)
	r.HandleFunc("/pg/countries", pgCountryHandler.All())
	r.HandleFunc("/pg/countries/{country}", pgCountryHandler.Uni())
	http.Handle("/", r)

	// Server
	log.Println("Starting server on ", os.Getenv("SERVER_ADDR"))
	srv := &http.Server{
		Handler:      r,
		Addr:         os.Getenv("SERVER_ADDR"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

// loadEnvVar loads environment variables from .env file
func loadEnvVar() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

// generatePgConnData generates PostgreSQL connection data
func generatePgConnData() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))
}
