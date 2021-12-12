all: gobuild docker compose

gobuild:
	GOOS=linux GOARCH=amd64 go build -o ./filemanagement/cmd/filemanagement-server/main ./filemanagement/cmd/filemanagement-server/main.go
timeout 5:
docker:
	 docker build -f ./build/Dockerfile -t filemgmt .
compose:
	 docker compose -f ./build/docker-compose.yml up -d
remove:
	 docker container stop filemgmtservice mongoservice
	yes |  docker-compose -f ./build/docker-compose.yml rm
	 docker rmi -f filemgmt
status:
	docker container ls -a 
	#############################################
	docker images
	#############################################
	docker-compose -f ./build/docker-compose.yml ps

swagger:
	swagger generate server -f ./filemanagement/api/filemgmt_swagger.yml -A filemanagement

	 
