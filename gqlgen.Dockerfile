FROM golang:1.18 as builder

WORKDIR /source

COPY go.mod go.mod
COPY go.sum go.sum
RUN GOCACHE=OFF
RUN go mod download

COPY cmd/gqlgen/gqlgen.yml cmd/gqlgen/gqlgen.yml
COPY cmd/gqlgen/server.go cmd/gqlgen/server.go
COPY cmd/gqlgen/graph cmd/gqlgen/graph/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o server cmd/gqlgen/server.go

FROM scratch
COPY .env .env
COPY --from=builder /source/server /server
CMD ["/server"]