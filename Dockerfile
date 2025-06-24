# Use official Golang image
FROM golang:1.22

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum first (for caching)
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the Go app
RUN go build -o smtp-server

# Expose the SMTP port
EXPOSE 2525

# Command to run the executable
CMD ["./smtp-server"]
