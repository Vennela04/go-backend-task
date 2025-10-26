# Build stage
FROM golang:1.25-alpine as builder
WORKDIR /app

# 🛠 Fix TLS issue: install certificates and enable TLS 1.2+
RUN apk add --no-cache ca-certificates openssl && update-ca-certificates

COPY go.mod go.sum ./
RUN GODEBUG=tls13=1 go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /go-backend ./cmd/server

# Final stage
FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /go-backend /go-backend
ENV DATABASE_URL=""
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/go-backend"]
