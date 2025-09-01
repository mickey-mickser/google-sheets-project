# Stage 1: Build the Go binary
FROM golang:1.23.6 AS build
WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download -x

# Copy all source and configs (service-account-key.json is excluded via .dockerignore)
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOARCH=amd64 go build -o /bin/docker-service ./cmd

# Stage 2: Create minimal runtime image
FROM alpine:latest AS final
WORKDIR /app

# Install CA certificates and timezone data
RUN apk update \
    && apk add --no-cache ca-certificates tzdata \
    && update-ca-certificates

# Copy the compiled binary
COPY --from=build /bin/docker-service /bin/docker-service

# Copy configs directory (dockerignore ensures no sensitive keys included)
COPY --from=build /app/configs /configs

# Copy .env for local development (ignored in production in favor of environment variables)
COPY --from=build /app/.env .env

# Copy JSON capabilities.json
COPY --from=build /app/capabilities.json /app/capabilities.json

# Create non-root user 'worker'
RUN adduser --disabled-password --gecos "" \
    --home "/nonexistent" --shell "/sbin/nologin" \
    --no-create-home --uid 10001 worker \
    && chown -R worker:worker /configs .env \
    && chmod 644 .env

# Switch to non-root user
USER worker

# Expose application port
EXPOSE 8080

# Entry point
ENTRYPOINT ["/bin/docker-service"]
