FROM golang:alpine as dev

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOD=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest as prod

WORKDIR /root

COPY --from=dev /app/main .
COPY --from=dev /app/.env .
COPY --from=dev /app/tables tables

EXPOSE 8080

ENTRYPOINT ["./main"]
