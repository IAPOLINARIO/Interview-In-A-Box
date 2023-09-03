# Build stage
FROM golang:1.20 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code to the workspace
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o slackviewer .

# Final stage
FROM alpine:latest

# Set the working directory
WORKDIR /app/

# Copy the slackviewer executable from the build stage to the final stage
COPY --from=builder /app/slackviewer .
COPY --from=builder /app/config.yml .

# Expose the application port
EXPOSE 8080

# Start the application
CMD ["./slackviewer", "server"]