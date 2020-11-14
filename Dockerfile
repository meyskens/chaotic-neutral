FROM golang:1.15-alpine as build

COPY ./ /go/src/github.com/meyskens/chaotic-neutral/
WORKDIR /go/src/github.com/meyskens/chaotic-neutral/

RUN go build ./

FROM alpine:3.12

COPY --from=build /go/src/github.com/meyskens/chaotic-neutral/chaotic-neutral /usr/local/bin

ENTRYPOINT ["chaotic-neutral"]