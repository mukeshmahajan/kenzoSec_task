# Step 1: Build the Go application
FROM golang:1.18-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies (this will be cached if the go.mod and go.sum files are not modified)
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the Go app (executable name will be 'app')
RUN go build -o app .

# Step 2: Create the final image with only the built application and necessary files
FROM alpine:latest  

# Set the working directory inside the container
WORKDIR /root/

# Copy the compiled Go binary from the builder image
COPY --from=builder /app/app .

# Copy the .env file for environment variables (if applicable)
COPY .env .

# Expose the port that the application will run on
EXPOSE 8080

# Command to run the application
CMD ["./app"]
