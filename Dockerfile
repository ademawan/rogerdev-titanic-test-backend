FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go build -o rogerdev-titanic-test-backend

EXPOSE 8081

CMD ./rogerdev-titanic-test-backend
