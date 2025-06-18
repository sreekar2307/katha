FROM golang:1.24-alpine AS builder
RUN apk add --no-cache make
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build

FROM alpine:latest
RUN apk add --no-cache make
WORKDIR /app
COPY --from=builder /app/bin/khata ./bin/khata
COPY Makefile ./
COPY config.yaml ./
EXPOSE 8080
CMD ["make", "http"]
