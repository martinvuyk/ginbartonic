FROM golang:1.19-alpine as build_stage

WORKDIR $GOPATH/src/

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@v1.8.12
RUN swag init

RUN go build -v -o /target-executable

FROM golang:1.19-alpine as deploy_stage
COPY --from=build_stage /target-executable $GOPATH/src/docs /