package main

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rangzen/isa-demo/pkg/graphql"
	"github.com/rangzen/isa-demo/pkg/handler"
	"github.com/rangzen/isa-demo/pkg/logs"
	"github.com/rangzen/isa-demo/pkg/pg"
	"github.com/rangzen/isa-demo/pkg/subscriber"
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
		log.Fatalf("opening PostgreSQL connection: %v", err)
	}
	defer func(pgDb *sql.DB) {
		err := pgDb.Close()
		if err != nil {
			log.Fatalf("closing the PostgreSQL database: %v", err)
		}
	}(pgDb)

	pgRepo := pg.NewRepository(pgDb)
	pgCountryHandler := handler.NewCountry(pgRepo)

	// Apollo
	apolloUrl := fmt.Sprintf("http://%s:%s%s", os.Getenv("APOLLO_HOST"), os.Getenv("APOLLO_PORT"), os.Getenv("APOLLO_PATH"))
	apolloRtr := graphql.NewRouter(apolloUrl)
	apolloEndPoint := handler.NewGraphQLEndpoint(apolloRtr)
	productQuery := "{\"query\": \"query { topProducts {name price}}\"}"

	// gqlgen
	gqlUrl := fmt.Sprintf("http://%s:%s%s", os.Getenv("GQLGEN_HOST"), os.Getenv("GQLGEN_PORT"), os.Getenv("GQLGEN_PATH"))
	gqlRtr := graphql.NewRouter(gqlUrl)
	gqlEndPoint := handler.NewGraphQLEndpoint(gqlRtr)
	todoQuery := "{\"query\": \"query findTodos {todos {text done user {name}}}\"}"

	// RedisPublish
	redisAddr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})
	if _, e := redisClient.Ping(redisClient.Context()).Result(); e != nil {
		log.Fatalf("pinging RedisPublish: %v", err)
	}

	// RedisPublish Publisher
	logRedisPublish := logs.NewRedisPublish(redisClient, os.Getenv("REDIS_LOG_CHANNEL"))

	// RedisPublish RedisSubscriber
	redisSubscribe := subscriber.NewRedisSubscriber(redisClient, os.Getenv("REDIS_LOG_CHANNEL"))
	go func() {
		if e := redisSubscribe.Subscribe(logs.Stdout); e != nil {
			log.Fatalf("subscribing to Redis Publisher: %v", err)
		}
	}()

	// Router
	r := mux.NewRouter()
	r.HandleFunc("/", handler.Home)
	// Middleware to logs to stdout
	r.HandleFunc("/pg/countries", logs.Handler(pgCountryHandler.All(), logs.Stdout))
	r.HandleFunc("/pg/countries/", pgCountryHandler.All())
	r.HandleFunc("/pg/countries/{country}", pgCountryHandler.Uni())
	// Middleware to logs to a RedisPublish publisher
	r.HandleFunc("/apollo/products", logs.Handler(apolloEndPoint.Handler(productQuery), logRedisPublish.Log))
	r.HandleFunc("/gql/todos", gqlEndPoint.Handler(todoQuery))
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
		log.Fatal("loading .env file")
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
