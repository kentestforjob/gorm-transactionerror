# for local development
FROM golang:1.19.5-alpine3.17

WORKDIR /app

COPY . .

RUN apk add --no-cache tzdata

RUN go mod download && go mod verify

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main . 

RUN go install github.com/cosmtrek/air@latest

CMD "air"

