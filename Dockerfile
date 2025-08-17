FROM golang:1.24.4 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o meuprimeirocrudgo .

FROM alpine:3.19 AS runner

RUN adduser -D -H -s /sbin/nologin jordan
WORKDIR /app
COPY --from=builder /app/meuprimeirocrudgo /app/meuprimeirocrudgo
RUN chown -R jordan:jordan /app && chmod +x /app/meuprimeirocrudgo

EXPOSE 8080
USER jordan
CMD ["./meuprimeirocrudgo"]
