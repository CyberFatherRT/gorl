FROM golang:1.22.2-bookworm AS builder
LABEL maintainer="CyberFatherRT"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go .
COPY pkg pkg
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o gorl .

FROM busybox:1.36.1-uclibc AS final
WORKDIR /app

COPY index.html .
COPY --from=builder /app/gorl gorl

ENTRYPOINT ["./gorl"]
