FROM golang:1.18 AS builder
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o server main.go

FROM alpine:latest
COPY --from=builder /src/server /server
ENTRYPOINT ["/server"]