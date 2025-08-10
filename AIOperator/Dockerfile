FROM golang:1.24.4 AS builder
WORKDIR /app
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main /app/src/api

FROM alpine:latest
ENV GIN_MODE=release
WORKDIR /app/
COPY --from=builder /main .
EXPOSE 8086
CMD ["/app/main"]
