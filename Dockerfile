FROM golang:1.12-alpine3.9

COPY . /go/src/github.com/hitman99/k8s-sandbox
WORKDIR /go/src/github.com/hitman99/k8s-sandbox
RUN go build -o sndbx

FROM alpine:3.9
LABEL maintainer="tomas@adomavicius.com"

RUN apk --no-cache add ca-certificates
WORKDIR /sandbox
COPY --from=0 /go/src/github.com/hitman99/k8s-sandbox/sndbx /sandbox/sndbx
COPY migrations /sandbox/migrations
ENV PATH="/sandbox/:${PATH}"

EXPOSE 8080
CMD ["sndbx"]