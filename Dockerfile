FROM alpine 

RUN mkdir /data
ADD ./hello /data
WORKDIR /data
EXPOSE 9090

ENTRYPOINT ["/data/hello"]  
