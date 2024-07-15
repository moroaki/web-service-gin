# Create a stage for building the application.
ARG GO_VERSION=1.22.5
FROM golang:${GO_VERSION}

# Install the Air CLI for live reloading during development.
RUN go install github.com/air-verse/air@latest
