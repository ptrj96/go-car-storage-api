FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -a -o /main

FROM alpine:latest
ARG PORT=8083
ENV APP_PORT=$PORT

WORKDIR /
COPY --from=builder /main /main
COPY listings.json /listings.json
EXPOSE $PORT

CMD ["/main"]
