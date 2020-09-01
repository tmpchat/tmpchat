FROM golang:1.14-alpine as build

WORKDIR /go/app

COPY . .

RUN apk add --no-cache git \
  && go build -o api_gateway \
  && go get -u github.com/cosmtrek/air


FROM alpine

WORKDIR /app

COPY --from=build /go/app/api_gateway .

CMD ["./api_gateway"]
