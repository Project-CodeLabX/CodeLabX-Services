FROM golang:1.21-alpine3.18 AS Builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd

EXPOSE 8010

FROM alpine:latest

WORKDIR /

COPY --from=Builder /api /api

EXPOSE 8010

CMD ["/api"]