FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go run github.com/beego/bee/v2@latest generate routers && \
    go run github.com/beego/bee/v2@latest generate docs
RUN CGO_ENABLED=0 go build -o server .

FROM alpine:3.19

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server .
COPY conf/app.conf conf/app.conf
COPY swagger/ swagger/

EXPOSE 8080

CMD ["./server"]
