FROM golang:1.21.1 AS build-server
LABEL stage=builder

WORKDIR /code/
# cache go modules
COPY go.mod /code/
COPY go.sum /code/
RUN go mod download

ADD .cache/upx /bin/upx
RUN chmod a+x /bin/upx

COPY ./ /code/
# CGO_ENABLED must be disabled for running in busyboxy/alpine
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /tmp/server .
RUN strip --strip-unneeded /tmp/server
RUN upx -q /tmp/server


FROM alpine:3.15 as alpine

RUN apk add -U --no-cache ca-certificates

FROM busybox AS runtime
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /home/dell/vsoutlook-app
COPY --from=build-server /tmp/server /home/dell/vsoutlook-app

ENV GIN_MODE release
CMD ["/home/dell/vsoutlook-app/server"]
