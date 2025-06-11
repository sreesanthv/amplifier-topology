# Use the official Golang image
FROM golang:1.24

# Set working directory
WORKDIR /app

# Copy Go source code
COPY . .

# Download dependencies
RUN go mod download

# Build the Go app
RUN go build -o nodes_topology

# Expose app port
EXPOSE 8080

# Run the app
CMD ["./nodes_topology"]