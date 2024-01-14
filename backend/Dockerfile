FROM golang:1.21.6-alpine
WORKDIR /backend
COPY go.mod  ./
RUN go mod download
COPY /src ./src
RUN go build -o ./dist/main src/main.go
CMD ["./dist/main"]
