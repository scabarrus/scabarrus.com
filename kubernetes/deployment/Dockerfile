FROM golang:alpine
RUN mkdir /app
COPY /root/go/src/scabarrus.com/kubernetes /app
WORKDIR /app/cmd/authn/
RUN go build -o main . 
CMD ["/app/cmd/authn/main"]
