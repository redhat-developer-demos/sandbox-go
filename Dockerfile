FROM docker.io/golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod init sandbox-go
#RUN go install github.com/gorilla/mux@latest
RUN go get github.com/gorilla/mux
RUN go get github.com/gin-gonic/gin
RUN go build -o  .
EXPOSE 8080
CMD ["/app/sandbox-go"]