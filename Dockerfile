# Imagen base para trabajar
FROM golang:alpine
RUN apk add git
ADD . /go/src/myapp
WORKDIR /go/src/myapp
RUN go get myapp
RUN go install
ENTRYPOINT ["/go/bin/myapp"]
