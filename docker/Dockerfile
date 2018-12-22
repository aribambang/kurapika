
#build stage
FROM golang:alpine AS builder
RUN mkdir -p /go/src/github.com/kurapika
COPY  . /go/src/github.com/kurapika
WORKDIR /go/src/github.com/kurapika
RUN go get ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kurapika .

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/kurapika/kurapika .
COPY ./.env .
RUN /bin/touch /root/kurapika.log