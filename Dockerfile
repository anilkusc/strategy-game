FROM golang:1.17.1 as build
RUN apt-get update && apt-get install sqlite3 -y
WORKDIR /src
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN go test -v ./...
RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o /bin/app .

FROM alpine
WORKDIR /app
COPY --from=build /bin/app .
CMD ["./app"]
#protoc --go_out=. --go-grpc_out=. -I=./api/ ./api/api.proto