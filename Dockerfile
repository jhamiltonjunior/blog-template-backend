FROM golang:1.17.7-buster


# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...


EXPOSE 3000