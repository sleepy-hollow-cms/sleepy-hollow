FROM golang:1.18 as buildImage

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /usr/src
COPY . .
RUN go build -o content-management-api main.go

# runtime image
FROM alpine:3.18.5

COPY --from=buildImage /usr/src/content-management-api /
COPY --from=buildImage /usr/src/conf/config.yaml /etc/sleepy-hollow/

ENTRYPOINT ["/content-management-api"]
