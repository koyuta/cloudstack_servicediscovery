FROM golang AS builder
LABEL maintainer "axm-misawa"

WORKDIR /go/src/app
COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN CGO_ENABLED=0 go build -o app ./

FROM busybox
COPY --from=builder /go/src/app/app /usr/local/bin/app
ENTRYPOINT ["/usr/local/bin/app"]
CMD ["--help"]
