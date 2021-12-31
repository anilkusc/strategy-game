FROM golang:1.17.1 as build
RUN apt-get update && apt-get install sqlite3 -y
WORKDIR /src
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN go test ./...
RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o /bin/app .

FROM scratch
WORKDIR /app
COPY --from=build /bin/app .
CMD ["app"]