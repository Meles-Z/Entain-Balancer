# Stage 1 — Build the Go binary
FROM golang:1.24-alpine AS builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go 

# Stage 2 — Run the binary in a lightweight image
FROM alpine:3.20

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/main .
# Only for development:
COPY .env .

EXPOSE 8080

CMD ["./main"]
