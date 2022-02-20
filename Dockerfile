FROM golang:1.17-buster
MAINTAINER robson mukubvu
COPY . /amakosi-login-service
WORKDIR /amakosi-login-service
RUN go build -o login-service
EXPOSE 8084
CMD ["./login-service"]

