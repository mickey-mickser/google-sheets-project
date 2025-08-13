#ARG GO_VERSION=1.23.6
#FROM golang:${GO_VERSION} AS build
#
#COPY . /github.com/mickey-mickser/google-sheets-poject/
#WORKDIR /github.com/mickey-mickser/google-sheets-poject/
#
#COPY go.mod go.sum ./
#RUN go mod download -x
#
#ARG TARGETARCH
#RUN CGO_ENABLED=0 GOARCH=amd64 go build -o /bin/bot ./cmd
#
#
#FROM alpine:latest AS final
#WORKDIR /root/
#
#COPY --from=build /bin/bot /bin/docker-service
#COPY --from=build /github.com/mickey-mickser/google-sheets-poject/configs /configs
#
#RUN apk update \
# && apk add --no-cache ca-certificates tzdata \
# && update-ca-certificates
#
#ARG UID=10001
#RUN adduser \
#    --disabled-password \
#    --gecos "" \
#    --home "/nonexistent" \
#    --shell "/sbin/nologin" \
#    --no-create-home \
#    --uid "${UID}" \
#    worker
#
#
#USER worker
#
#EXPOSE 8080
#
#ENTRYPOINT ["/bin/docker-service"]
########################################
########################################
# Build stage
########################################
#FROM golang:1.23.6 AS build
#
#WORKDIR /app
#
## Copy go modules manifests
#COPY go.mod go.sum ./
#RUN go mod download -x
#
## Copy source code and .env for godotenv
#COPY . .
#
## Build the Go binary
#RUN CGO_ENABLED=0 GOARCH=amd64 go build -o /bin/docker-service ./cmd
#
#########################################
## Final stage
#########################################
#FROM alpine:latest AS final
#
#WORKDIR /app
#
## Install certificates and timezone data
#RUN apk update \
#    && apk add --no-cache ca-certificates tzdata \
#    && update-ca-certificates
#
## Copy built binary
#COPY --from=build /bin/docker-service /bin/docker-service
#
## Copy configs and .env
#COPY --from=build /app/configs /configs
#COPY --from=build /app/.env .env
#
## --- ИСПРАВЛЕНИЕ: Перемещаем создание пользователя ПЕРЕД chown ---
## Создаем не-root пользователя 'worker'
## --disabled-password: не устанавливать пароль
## --gecos "": не запрашивать информацию о пользователе
## --home "/nonexistent": не создавать домашнюю директорию
## --shell "/sbin/nologin": запретить вход через оболочку
## --no-create-home: явно не создавать домашнюю директорию
## --uid 10001: установить конкретный UID
#RUN adduser --disabled-password --gecos "" --home "/nonexistent" --shell "/sbin/nologin" --no-create-home --uid 10001 worker
#
## Убедитесь, что .env доступен для чтения не-root пользователем
#RUN chown worker:worker .env \
#    && chmod 644 .env
#
## Переключаемся на не-root пользователя
#USER worker
#
## Expose application port
#EXPOSE 8080
#
## Entrypoint
#ENTRYPOINT ["/bin/docker-service"]
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
