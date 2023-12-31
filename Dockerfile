FROM golang:latest
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api/main.go
CMD ["./api"]

