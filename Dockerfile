FROM golang:alpine AS builder
LABEL	maintainer "axm-misawa"

WORKDIR /go/src/app
COPY . .

RUN go get -u github.com/golang/dep/...
RUN dep ensure
RUN go build -o app main.go

FROM busybox
COPY --from=builder /go/src/app/app /usr/local/bin/app
ENTRYPOINT ["/usr/local/bin/app"]
