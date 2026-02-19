# Build Stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git
COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux \
    go build -ldflags="-s -w" -o app

# Runtime Stage
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/app/app" ]