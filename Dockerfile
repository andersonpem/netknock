# Use the official Go image based on Musl for cross-compilation
FROM --platform=$TARGETPLATFORM golang:1.20 AS builder

# Set the working directory
WORKDIR /go/src/app

# Copy the Go source code into the image
COPY . .
WORKDIR /go/src/app/src

# We must fetch dependencies
RUN go get .

# Install Musl tools for cross-compilation
RUN apt-get update && apt-get install -y musl-tools

# Build the Go app for the target architecture (e.g., amd64)
RUN CC=musl-gcc GOARCH=$TARGETARCH CGO_ENABLED=1 go build -o netknock main.go

# Use a minimal Alpine image for the runtime
FROM --platform=$TARGETPLATFORM alpine:latest
LABEL org.opencontainers.image.authors="AndersonPEM https://github.com/andersonpem"

# Copy the "netknock" binary from the builder stage to /usr/local/bin
COPY --from=builder /go/src/app/src/netknock /usr/local/bin/netknock

# Make the binary executable
RUN chmod +x /usr/local/bin/netknock

# Run the binary
ENTRYPOINT ["/usr/local/bin/netknock"]
