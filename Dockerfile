FROM golang:1.17 as builder
WORKDIR /go/src/sample_exporter_go
COPY . /go/src/sample_exporter_go
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata
WORKDIR /work
COPY --from=builder /go/src/sample_exporter_go/main /work
ENTRYPOINT ["./main"]
