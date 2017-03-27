FROM alpine 

RUN mkdir /data
ADD ./hello /data
WORKDIR /data
EXPOSE 8080

ENTRYPOINT ["/data/hello"]  
