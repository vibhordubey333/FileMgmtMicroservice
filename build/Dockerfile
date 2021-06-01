FROM golang:1.14

MAINTAINER vibhordubey333@gmail.com

#VOLUME .cmd/filemanagement-server/:/app/

WORKDIR /app

COPY ./filemanagement/cmd/filemanagement-server/main /app/

ENTRYPOINT ["/app/main"]