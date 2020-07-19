FROM golang:1.14-alpine as build

WORKDIR /go/app

COPY . .

RUN apk add --no-cache git \
  && go build -o message_broker


FROM alpine

WORKDIR /app

COPY --from=build /go/app/message_broker .

CMD ["./message_broker"]
