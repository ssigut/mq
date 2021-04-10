# Build image
FROM golang:1.15 as build
ENV CGO_ENABLED=0

COPY ./ /src/mq/
WORKDIR /src/mq
RUN mkdir /out

RUN GOOS=linux GOARCH=amd64 go build -o /out/mq main.go

# Runtime image
FROM alpine
RUN apk add --update ca-certificates && \
    rm -rf /var/cache/apk/* /tmp/*
COPY --from=build /out/mq /
WORKDIR /

ENTRYPOINT ["/mq"]
