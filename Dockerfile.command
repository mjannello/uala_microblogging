FROM golang:1.21

WORKDIR /app

COPY go.mod .

RUN go mod tidy
RUN go mod download

COPY . .

RUN go build -o command-app cmd/command/main.go

CMD ["./command-app"]
