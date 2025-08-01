FROM golang:1.24.2-alpine3.11 AS build

RUN apk --no-cache add gcc g++ make ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
COPY vendor ./vendor

COPY catalog ./catalog

RUN go build -mod=vendor -o /app/catalog-service ./catalog/cmd

FROM alpine:3.11

WORKDIR /usr/bin

COPY --from=build /app/catalog-service ./app
EXPOSE 8080

CMD ["./app"]
