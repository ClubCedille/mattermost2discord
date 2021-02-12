# Step 1 - compile source code
FROM golang:1.15.6-buster AS builder

LABEL maintainer="Club CEDILLE"

ADD . /go/src/app
WORKDIR /go/src/app

RUN go get app
RUN go install
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build -a -installsuffix cgo -o mm2disc .

# Step 2 - download ca-certs
FROM alpine:latest as certs
RUN apk --update add ca-certificates

# Step 3 - run the binary
FROM scratch
COPY --from=builder /go/src/app/mm2disc .
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT ["/mm2disc"]
