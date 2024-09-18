# Start from official GO image
FROM golang:1.18

# Set up the working directory
WORKDIR /app

# Copy the Go module and sum files
COPY go.mod ./
COPY go.sum ./

# Install dependencies
RUN go mod download

# Copy the src code
COPY . .

# Build the Go Server
RUN go build -o ./grpc_server ./main.go

# Expose port for the gRPC server
EXPOSE 50051

# Run the gRPC server
CMD ["./grpc_server"]