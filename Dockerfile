# Step 1 - compile source code
FROM golang:1.16.4-buster AS builder

LABEL maintainer="Club CEDILLE"

WORKDIR /app
COPY . /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build -a -installsuffix cgo -o mm2disc .

# Add user & group
RUN groupadd mm2disc && \
    useradd -r -u 1001 -g mm2disc mm2disc

# Step 2 - download ca-certs
FROM alpine:latest as certs
RUN apk --update add ca-certificates

# Step 3 - run the binary
FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app/mm2disc .
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Set which user to use
USER mm2disc

ENTRYPOINT ["/mm2disc"]
