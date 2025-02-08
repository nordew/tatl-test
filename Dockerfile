FROM golang:1.23.4-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY .env .

RUN go mod download
RUN go mod tidy

COPY . .

RUN go build -o app ./cmd/app

CMD ["./app"]
