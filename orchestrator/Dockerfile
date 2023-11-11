FROM golang:1.18 AS builder
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o orchestrator main.go

FROM alpine:latest
COPY --from=builder /src/orchestrator /orchestrator
ENTRYPOINT ["/orchestrator"]