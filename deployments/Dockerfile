FROM golang:1.23-alpine as builder

WORKDIR /app

COPY /go.mod /go.sum ./

RUN go mod download

COPY . .

COPY ../config/local.env /app/config/local.env

RUN CGO_ENABLED=0 GOOS=linux go build -o delivery_backend ./cmd/app

CMD ["/delivery_backend"]

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/delivery_backend .

COPY --from=builder /app/config/local.env /root/config/local.env

EXPOSE 8080

CMD ["./delivery_backend"]