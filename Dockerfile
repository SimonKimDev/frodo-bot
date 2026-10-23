# syntax=docker/dockerfile:1

FROM golang:1.26-alpine AS build
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/frodo-bot ./cmd/frodo-bot

# distroless/static has CA certificates (needed for Discord's TLS) and nothing else; runs as a non-root user.
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/frodo-bot /frodo-bot
ENTRYPOINT ["/frodo-bot"]
