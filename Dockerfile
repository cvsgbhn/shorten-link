# syntax=docker/dockerfile:1

FROM golang:1.13-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -o docker-app cmd/app/main.go 
RUN ls -al

ARG db

ENV DBType=$db

RUN echo $DBType

#ENTRYPOINT ["sh", "./docker-params.sh"]
ENTRYPOINT ./docker-app $DBType