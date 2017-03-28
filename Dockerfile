FROM alpine 

RUN mkdir -p /data
ADD hello /data
WORKDIR /data

EXPOSE 8080

ENTRYPOINT ["/data/hello"]
