FROM golang:1.14.0-alpine3.11 AS build
RUN apk --no-cache add gcc g++ make ca-certificates git
WORKDIR /go/src/github.com/hirac1220/gcp-test/gae
RUN go build -o /go/bin/go-hello

FROM alpine:3.11
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["go-hello"]