FROM golang:alpine

WORKDIR /go/src/app
COPY ./main.go .

EXPOSE  8081

RUN go get -d -v ./...
RUN go install -v ./...

CMD [ "app" ]