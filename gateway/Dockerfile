FROM golang:1.10.1-alpine3.7 as builder
COPY . /go/src/graphql/
WORKDIR /go/src/graphql
RUN apk update \
  && apk add --no-cache git \
  && go get -u github.com/golang/dep/cmd/dep \
  && dep ensure \
  && go build -o /app main.go

FROM alpine:3.7
CMD ["./app"]
COPY --from=builder /app .
