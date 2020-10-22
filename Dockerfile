FROM golang:1.14.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags "-X main.githash=$(git rev-parse HEAD) -X main.buildstamp=$(date +%Y%m%d.%H%M%S)" \
    -o goapp .

################################################

FROM alpine:latest

RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime && \
    echo "Asia/Bangkok" >/etc/timezone && \
    apk del tzdata

WORKDIR /app

COPY --from=builder /app/config.yaml .
COPY --from=builder /app/goapp .

CMD ["/app/goapp"]
