FROM alpine:3.9
COPY ./bin/gin-web-demo .
COPY ./application.yaml .
EXPOSE 8188
ENTRYPOINT ["/gin-web-demo"]
