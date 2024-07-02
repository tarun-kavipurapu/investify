FROM golang:1.22.2-alpine3.19 AS builder
WORKDIR /

COPY go.mod go.sum ./

COPY . .
RUN apk add --no-cache git

RUN go mod download

RUN go build -o cmd/app/main .

EXPOSE 8000

CMD [ "./main" ]


