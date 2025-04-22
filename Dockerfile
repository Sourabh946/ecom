# Start from the official Go image
FROM golang:1.22

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./

# Download Go modules
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main .

# Run the binary executable
CMD ["./main"]
