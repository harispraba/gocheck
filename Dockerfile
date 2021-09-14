FROM golang:1.16.8-alpine3.14 as build

ARG VERSION

LABEL description="Golang SSL Checker"
LABEL repository="https://github.com/mrofisr/gocheck"
LABEL maintainer="mrofisr"

WORKDIR /app
COPY ./go.mod .
RUN go mod download

COPY . .
RUN go build .

FROM alpine:latest

COPY --from=build /app/gocheck /bin/gocheck
ENV HOME /
ENTRYPOINT ["/bin/gocheck"]
