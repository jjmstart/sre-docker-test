FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/myapp .
CMD ["./myapp"]
