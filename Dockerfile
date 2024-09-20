FROM golang:1.23.1 AS builder

WORKDIR /app

# This is so the server can be accessed from outside of the container
# Check out .env.example for other values
ENV APP_SERVER_HOST=0.0.0.0

# Copy and cache the dependencies first
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the project files
COPY . .

# Package and build the project into an executable binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o main -v -trimpath -x ./cmd/server/

# Switch to a smaller image
FROM alpine:latest

WORKDIR /app

# Add a non-root user and group
RUN addgroup -S goapp && adduser -S gouser -G goapp

# Copy the binary from the previous stage to the current image
COPY --from=builder /app/main .

# Change the owner of the file to the newly created user
RUN chown gouser:goapp /app/main

# Switch to the new user
USER gouser

# Run the server
CMD ["./main"]
