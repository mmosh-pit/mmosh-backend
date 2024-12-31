FROM golang:latest AS builder
WORKDIR /builder
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mmosh-backend ./cmd

FROM alpine:latest AS app
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /builder/mmosh-backend ./mmosh-backend
COPY --from=builder /builder/.env ./.env
EXPOSE 8000
CMD ["./mmosh-backend"]
