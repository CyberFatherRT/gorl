FROM golang:1.22.2-bookworm AS builder
WORKDIR /app

COPY . .
RUN go build -o gorl main.go

FROM scratch AS final
WORKDIR /app

COPY --from=builder /app/gorl gorl

ENTRYPOINT ["/app/gorl"]
