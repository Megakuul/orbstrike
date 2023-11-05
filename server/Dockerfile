FROM golang:1.18 AS builder
WORKDIR /src
COPY . .
RUN go build -o /server main.go

FROM alpine:latest
COPY --from=builder /server /server
ENTRYPOINT ["/server"]