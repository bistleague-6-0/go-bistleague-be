FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary application/rest/*.go

EXPOSE 3381

ENTRYPOINT ["/app/binary"]