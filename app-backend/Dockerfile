FROM golang:1.13.7-alpine3.11

WORKDIR /tmp/buildDir
COPY ./src src
RUN go build -o be-app ./src

FROM alpine:3.11
WORKDIR /opt/creeptoo
COPY --from=0 /tmp/buildDir/be-app .
CMD ["./be-app"]