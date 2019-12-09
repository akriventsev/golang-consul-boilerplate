# Start from a Alpine image
FROM alpine:latest
# Copy the local package files to the container's workspace.
ADD ./bin /go/bin
# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/service-linux-amd64
