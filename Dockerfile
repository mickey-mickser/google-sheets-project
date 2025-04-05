ARG GO_VERSION=1.23.6
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build

COPY . /github.com/mickey-mickser/google-sheets-poject/
WORKDIR /github.com/mickey-mickser/google-sheets-poject/

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

ARG TARGETARCH
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/bot ./cmd

FROM alpine:latest AS final
WORKDIR /root/

COPY --from=build /bin/bot /bin/docker-service
COPY --from=build /github.com/mickey-mickser/google-sheets-poject/configs /configs

RUN --mount=type=cache,target=/var/cache/apk \
    apk --update add \
        ca-certificates \
        tzdata \
        && \
        update-ca-certificates

ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    worker


USER worker

EXPOSE 8055

ENTRYPOINT ["/bin/docker-service"]
