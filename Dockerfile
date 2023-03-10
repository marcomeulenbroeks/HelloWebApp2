FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN go build -o /hello-web-app

EXPOSE 8087

CMD [ "/hello-web-app" ]