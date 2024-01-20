FROM golang:1.21.6-alpine
WORKDIR /seed
COPY go.mod  ./
COPY go.sum ./
RUN go mod download
COPY database/ ./database
COPY src/database ./src/database
COPY src/logger ./src/logger
RUN go build -o ./dist/main ./database/mongo_seed/main.go
CMD ["./dist/main"]
