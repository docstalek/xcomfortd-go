FROM golang:alpine3.20 as builder

RUN apk add --no-cache --update gcc musl-dev linux-headers pkgconfig libusb-dev

COPY . /xcomfortd-go

WORKDIR /xcomfortd-go
RUN go get
RUN go install

FROM alpine:3.20

RUN apk add --no-cache --update libusb
COPY --from=builder /go/bin/xcomfortd-go /usr/bin/xcomfortd

CMD ["xcomfortd"]
