FROM golang:1.18 AS builder
WORKDIR /src
COPY . .
RUN go build -o /orchestrator main.go

FROM alpine:latest
COPY --from=builder /orchestrator /orchestrator
ENTRYPOINT ["/orchestrator"]