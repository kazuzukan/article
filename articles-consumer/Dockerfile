FROM golang:1.24.1-alpine

# Set working directory
WORKDIR /app

# Copy go mod/sum and download dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN go build -o main .

# Run the binary
CMD ["./main"]
