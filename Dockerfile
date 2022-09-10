# build stage
FROM golang:1.19-alpine AS builder

LABEL maintainer="Joel Santos <joe@joesantos.io>"

WORKDIR /app
COPY . .

RUN apk update && apk add --virtual build-dependencies build-base gcc git make

RUN go mod download
RUN go build -v -o ./app

# final stage
FROM alpine

ENV ENV=production
ENV PORT=4040
ENV SERVER_PACKAGE=chi
ENV DB_TYPE=inmem

WORKDIR /root/
COPY --from=builder /app/app ./app

CMD ["./app"]