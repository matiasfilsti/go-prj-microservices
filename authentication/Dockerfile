# Usar una imagen base para la construcción
# FROM golang:1.23.2 AS builder
# WORKDIR /app
# COPY ./go.mod ./go.sum ./
# RUN go mod download
# COPY src/api/ .
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main .

FROM golang:1.23.2 AS builder
WORKDIR /app
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main /app/src/api

FROM alpine:latest
ENV GIN_MODE=release
WORKDIR /app/
COPY --from=builder /main .
EXPOSE 8080
CMD ["/app/main"]
