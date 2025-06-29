FROM golang:1.23.2 AS builder
WORKDIR /app
COPY . ./
RUN go mod download
RUN cd src/api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main /app/src/api

FROM alpine:latest
ENV GIN_MODE=release
WORKDIR /app/
COPY --from=builder /main .
COPY --from=builder /app/html-templates ./html-templates
EXPOSE 8085
CMD ["/app/main"]
