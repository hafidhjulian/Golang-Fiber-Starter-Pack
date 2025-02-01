FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git openssh openssh-client docker docker-compose openrc sudo nano curl tzdata

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

FROM alpine:latest AS runner

RUN apk --no-cache add tzdata

WORKDIR /app

COPY --from=builder /app .

ENTRYPOINT ["/app/binary"]