FROM nicolaka/netshoot:latest
ADD ./bin/grpc-server /data/grpc-server

WORKDIR /data
CMD ["./grpc-server"]
