FROM golang:1.22.2-alpine as builder

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o nearby ./cmd/api

RUN chmod +x /app/nearby

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/nearby /app

CMD ["/app/nearby"]