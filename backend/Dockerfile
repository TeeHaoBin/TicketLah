FROM golang:1.25.6-alpine

WORKDIR /src/app

# Copy go mod files first for cache efficiency
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build binary at image build time (avoids virtiofs deadlock on macOS)
RUN go build -o /usr/local/bin/app ./cmd/api

CMD ["/usr/local/bin/app"]