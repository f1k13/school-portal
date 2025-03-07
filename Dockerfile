FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM debian:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["/app/main"]
