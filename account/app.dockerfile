FROM golang:1.24.2-alpine3.18 AS build

RUN apk --no-cache add gcc g++ make ca-certificates

# Set the working directory inside the container
WORKDIR /app

COPY go.mod go.sum ./
COPY vendor vendor

COPY account account

RUN go build -mod=vendor -o /app/account-service ./account/cmd

FROM alpine:3.18

# Set working directory
WORKDIR /usr/bin

# Copy the compiled binary from the builder stage
COPY --from=build /app/account-service ./app
EXPOSE 8080

# Command to run the binary
CMD ["./app"]
