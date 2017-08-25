FROM golang:alpine AS build-env
WORKDIR /usr/local/go/src/github.com/mattn/lile-example/gene9go
COPY . /usr/local/go/src/github.com/mattn/lile-example/gene9go
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get ./...
RUN go build -o build/gene9go ./gene9go


FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=build-env /usr/local/go/src/github.com/mattn/lile-example/gene9go/build/gene9go /bin/gene9go
CMD ["gene9go", "up"]
