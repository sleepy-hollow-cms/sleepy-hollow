FROM golang:1.17 as buildImage

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /usr/src
COPY . .
RUN go build -o content-management-api main.go

# runtime image
FROM alpine:3.14.2

COPY --from=buildImage /usr/src/content-management-api /
COPY --from=buildImage /usr/src/conf/config.yaml /etc/sleepy-hollow/

ENTRYPOINT ["/content-management-api"]
