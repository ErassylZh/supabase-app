FROM docker-registry.fmobile.kz/golang:1.19-alpine

RUN ln -snf /usr/share/zoneinfo/Asia/Almaty /etc/localtime && echo Asia/Almaty > /etc/timezone

RUN apk update && apk upgrade && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o index

RUN chmod -R 777 /app/index

RUN adduser -D -g 'app' app -u 1001

USER app

WORKDIR /app

ENTRYPOINT ["/app/index"]
