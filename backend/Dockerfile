FROM golang:1.25.6-alpine

WORKDIR /src/app

# Install air (pinned version)
RUN go install github.com/air-verse/air@latest

# Copy go mod files first for cache efficiency
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# air will run via docker-compose