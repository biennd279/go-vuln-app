# Base image
FROM golang:1.22.0-alpine3.19 AS builder

# Author
LABEL maintainer="biennd5"

# Set the Current Working Directory inside the container
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/web/main.go

FROM alpine:3.19 AS production

WORKDIR /app

COPY --from=builder /app/main  .

ENV GIN_MODE=release

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

