# ISA Demo

Yes, it is a geek joke: [ISA](https://en.wikipedia.org/wiki/Industry_Standard_Architecture).

## General

Backend with PostgreSQL, Apollo Server and gqlgen.

## Points of interest

* No preemptive interface (e.g., [CountryRepository in country.go](pkg/handler/country.go)).
* Tests from external package (e.g., [country_test.go](pkg/handler/country_test.go)).
* Docker for gqlgen [gqlgen.Dockerfile](gqlgen.Dockerfile).
* Docker-compose for the whole stack [docker-compose.yml](docker-compose.yml).
* [/docs](docs)
  * [api.http](docs/api.http) for IntelliJ REST Client.
* `.env` file for configuration.

## Target

### SQL

Using Jet to get info with a join between two tables.

* [x] [PostgreSQL](https://www.postgresql.org/)
* [x] [Jet](https://github.com/go-jet/jet)

### GraphQL

Using gqlgen to access to a view of a data table.

* [x] [Apollo](https://www.apollographql.com/)
* [x] [gglgen](https://gqlgen.com/)

### Redis

Cache with automatic ttl some results.

* [ ] [Redis](https://redis.io/)

### Heroku

* [ ] Having everything online.

### Camunda

?
