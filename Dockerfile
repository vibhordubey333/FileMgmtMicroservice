FROM golang:1.14

#ADD ./cmd/file-mgmt-microservice-server/main /

ADD ./cmd/file-mgmt-microservice-server/main /

ENTRYPOINT ["/main"]
