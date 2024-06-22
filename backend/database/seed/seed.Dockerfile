FROM golang:1.21.6-alpine
WORKDIR /seed
COPY go.mod  ./
COPY go.sum ./
RUN go mod download
COPY database/ ./database
COPY src/database ./src/database
COPY src/logger ./src/logger
COPY src/events ./src/events
COPY src/helper/auth0.go ./src/helper/auth0.go
RUN go build -o ./dist/main ./database/seed/main.go