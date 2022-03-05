FROM golang

# ADD . /go/src/github.com/programadriano/golang-api/
# COPY . /go/src/github.com/programadriano/golang-api/

# RUN go get github.com/gorilla/mux
# RUN go install github.com/programadriano/golang-api

# ENTRYPOINT /go/bin/golang-api

EXPOSE 3000