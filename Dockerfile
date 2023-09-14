# Use the official Go image to build the binary
FROM --platform=$TARGETPLATFORM golang:1.20 AS builder

# Set the working directory
WORKDIR /go/src/app

# Copy the Go source code into the image
COPY src/* .

# Build the Go app and name the binary "netknock"
RUN go build -o netknock main.go

# Use a minimal alpine image for the runtime with architecture in mind
FROM --platform=$TARGETPLATFORM alpine:latest
LABEL org.opencontainers.image.authors="AndersonPEM https://github.com/andersonpem"

# Copy the "tcpknock" binary from the builder stage to /usr/local/bin
COPY --from=builder /go/src/app/netknock /usr/local/bin/netknock

# Make the binary executable
RUN chmod +x /usr/local/bin/netknock

# Run the binary
ENTRYPOINT ["/usr/local/bin/tcpknock"]
