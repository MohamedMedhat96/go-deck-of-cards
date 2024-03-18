
FROM golang:1.18 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /app/bin/app ./cmd/api

FROM alpine:3
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/bin/app /app/bin/app

ENV SERVER_PORT=8080 \
    MONGO_DB_URI=mongodb://mongodb:27017 \
    DATABASE_NAME=deck \
    SERVER_HOST=0.0.0.0

EXPOSE $SERVER_PORT

CMD ["/app/bin/app"]
