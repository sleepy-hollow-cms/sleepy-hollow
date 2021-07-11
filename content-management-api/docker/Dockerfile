FROM golang:1.16 as buildImage

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /usr/src
COPY . .
RUN go build -o content-management-api main.go

# runtime image
FROM alpine:3.14.0

COPY --from=buildImage /usr/src/content-management-api /

ENTRYPOINT ["/content-management-api"]
