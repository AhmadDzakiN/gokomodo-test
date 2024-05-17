FROM golang:1.21-alpine3.19

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

# To configure your env value, please check your params/.env file and edit it first before creating the app's docker container

COPY params/.env .

RUN go mod tidy

RUN go build -o /gokomodo-assignment

RUN chmod +x /gokomodo-assignment

EXPOSE 8000

CMD ["/gokomodo-assignment"]