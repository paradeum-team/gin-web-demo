FROM centos:7
COPY ./bin/gin-web-demo /data/
COPY ./application.yaml /data/
EXPOSE 8188
WORKDIR /data

CMD ["./gin-web-demo"]
