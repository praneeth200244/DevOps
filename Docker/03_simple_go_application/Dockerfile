# Step 1: Build the Go app
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum (if it exists)
COPY go.mod ./

# Copy the Go source code into the container
COPY . .

# Build the Go app
RUN go build -o welcome-app main.go

# Step 2: Create a smaller image for running the Go app
FROM alpine:latest

# Install dependencies to run the Go app (e.g., for terminal input/output)
RUN apk --no-cache add libc6-compat

# Set the working directory
WORKDIR /app

# Copy the built Go binary from the builder stage
COPY --from=builder /app/welcome-app /app/welcome-app

# Run the Go binary when the container starts
ENTRYPOINT ["./welcome-app"]
