FROM golang:1.13.1
RUN go get -d -v github.com/leros321/quay-exporter/quay_exporter
WORKDIR /go/src/github.com/leros321/quay-exporter/quay-exporter
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o quay-exporter .

FROM alpine:3.7
WORKDIR /root/
LABEL maintainer="damien@weave.works"
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=0 /go/src/github.com/leros321/quay-exporter/quay-exporter/quay-exporter .
EXPOSE 8080
ENTRYPOINT ["/usr/bin/quay-exporter"]
CMD ["-help"]
