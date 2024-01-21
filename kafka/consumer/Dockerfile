FROM bitnami/golang AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN GOPROXY=direct go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o consumer .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/consumer .
CMD ["./consumer"]