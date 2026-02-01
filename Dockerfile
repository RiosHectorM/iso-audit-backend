# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./
# COPY go.sum ./ # go.sum might not exist yet

RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main ./cmd/api/main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
