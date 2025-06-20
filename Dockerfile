FROM golang:1.24-alpine

WORKDIR /app

# Install Air
RUN go install github.com/air-verse/air@latest

# Cache go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy code
COPY . .

# Start app with Air
CMD ["air", "-c", ".air.toml"]