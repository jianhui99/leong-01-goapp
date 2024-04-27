FROM golang:1.20
WORKDIR /go/src
ADD . /go/src
RUN cd /go/src 
RUN go mod tidy
RUN go build -o main
CMD ["./main"]