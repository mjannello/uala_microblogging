FROM golang:1.21

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o uala_app cmd/command/main.go

CMD ["./uala_app"]
