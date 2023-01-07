FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy

CMD CGO_ENABLED=0 go test --tags=unit -v ./...

RUN go build -o expenses-tracking .

FROM alpine:3.16.3

EXPOSE 2565

WORKDIR /app

COPY --from=builder /app/expenses-tracking /app

CMD [ "/app/expenses-tracking" ]