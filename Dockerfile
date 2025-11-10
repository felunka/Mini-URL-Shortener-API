# Build Stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .

RUN go mod tidy && go build -o url-shortener

# Final Stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/url-shortener .
CMD ["./url-shortener"]

EXPOSE 8080
