# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Sebastien Cabarrus <scabarrus@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /go/src

# Copy the source from the current directory to the Working Directory inside the container
COPY ./scabarrus.com ./scabarrus.com

RUN cd scabarrus.com/k8s.webhook ; go mod init; go mod vendor;go get -u github.com/swaggo/swag/cmd/swag

#Download swag command and generate swagger with dependencies
RUN swag init --parseDependency -g scabarrus.com/k8s.webhook/cmd/user-management/main.go 

# Build the Go app
RUN go build -o main  scabarrus.com/k8s.webhook/cmd/user-management


# Expose port 9443 to the outside world
EXPOSE 9443

# Command to run the executable
CMD ["./main"]
