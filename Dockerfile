FROM golang:alpine 

RUN mkdir /data
ADD hello-app /data
RUN cd /data/hello-app
RUN go build -o hello .
WORKDIR /data/hello-app
EXPOSE 8080

ENTRYPOINT ["/data/hello-app/hello"]  
