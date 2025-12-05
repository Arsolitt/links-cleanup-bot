FROM golang:1.25.4-bookworm AS builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o .build/app .

FROM debian:bookworm-slim AS runtime
ENV TZ=UTC
WORKDIR /app
COPY --from=builder /src/app/.build/app .
CMD ["./app"]
