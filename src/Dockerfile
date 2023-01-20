FROM golang:1.19-alpine as build_stage

ENV APP_NAME api
WORKDIR $GOPATH/src/$APP_NAME

COPY go.mod go.sum ./
RUN go mod download

COPY . .
 
RUN go build -v -o /target-executable

FROM golang:1.19-alpine as deploy_stage
COPY --from=build_stage /target-executable /