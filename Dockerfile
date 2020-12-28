FROM node:15.5.0-alpine3.12 AS ASSETS

RUN mkdir -p /app
WORKDIR /app

COPY package.json .
COPY package-lock.json .
COPY resources resources
COPY public public

RUN npm i
RUN npm run prod

FROM golang:1.15.6-alpine3.12 AS BUILD

RUN mkdir -p /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go install

COPY . .

RUN go build -o oq main.go

FROM alpine:3.12.3 AS EXPORT
LABEL maintainer="Amir Ghaffari<contact@amirghaffari.com>"
RUN mkdir -p /app/resources
WORKDIR /app
EXPOSE 8080

COPY --from=ASSETS public public
COPY resource/template resource/template
COPY --from=build oq oq
COPY data data

CMD ./oq