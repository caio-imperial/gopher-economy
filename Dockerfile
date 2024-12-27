FROM golang:1.22.10-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o gopher-economy

FROM alpine

WORKDIR /app

RUN apk add --no-cache ca-certificates curl

COPY --from=builder /app .

ENTRYPOINT ["./gopher-economy"]