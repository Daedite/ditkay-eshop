FROM golang:1.19-alpine
RUN apk update && apk add --no-cache git
RUN apk add build-base
RUN apk --no-cach add ca-certificates && update-ca-certificates

ENV CGO_ENABLED 0

WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
RUN go mod verify

COPY . .
RUN cd cmd/dit-kay-eshop-api && go build

FROM scratch

EXPOSE 3001

WORKDIR /app
COPY --from=0 /app/cmd/dit-kay-eshop-api/dit-kay-eshop-api usr/local/bin/dit-kay-eshop-api
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ADD db/migrations/* /app/db/migrations/
CMD ["dit-kay-eshop-api"]
