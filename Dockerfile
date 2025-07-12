# Use Go image
FROM golang:1.22

# Set working directory
WORKDIR /app

# Copy files
COPY . .

# Download dependencies
RUN go mod download

# Build producer
RUN go build -o producer producer.go

# Run the app
CMD ["./producer"]
