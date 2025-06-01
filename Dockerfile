FROM golang:1.24.3-bullseye AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
# go関連のファイルが入る
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .


EXPOSE 8080
CMD ["/app/main"]