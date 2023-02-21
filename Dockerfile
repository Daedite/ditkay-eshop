FROM golang:latest

LABEL maintainer="Espoir Ditekemena <espopir@ditkay.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 3001

ADD db/migration/* /app/db/migration/

CMD ["./main"]