# Step 1 builder
FROM golang:1.15-alpine3.13 AS builder
RUN apk update && apk add --no-cache git make
WORKDIR /home
COPY . .
RUN make bin_api

# Step 2 build image
FROM alpine:3.13

RUN apk update && apk add --no-cache curl ca-certificates
RUN rm -rf /var/cache/apk/*

COPY --from=builder /home/waitress .

EXPOSE 8080

CMD ./waitress
