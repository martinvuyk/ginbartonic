FROM golang:1.19-alpine

ENV APP_NAME microservice3
 
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME
 
RUN go build -v -o /target-executable