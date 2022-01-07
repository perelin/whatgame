FROM golang:1.16-alpine as builder

## adding general build stuff -> seems to be neede to run `go test` a container

RUN apk update  \
    && apk add --virtual build-dependencies  \
    build-base  \
    gcc  \
    git  \
    curl \
    wget  \
    && rm -rf /var/cache/apk/*

## app

USER root
RUN mkdir /build 
ADD ./ /build/
WORKDIR /build 
RUN go build -v -o server

### 
## production image 
### 

FROM alpine

USER appuser
COPY --from=builder /build/server  /app/server
CMD ["/app/server"]