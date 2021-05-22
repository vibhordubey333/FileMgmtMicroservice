all: docker compose

docker:
	 docker build -f build/Dockerfile -t filemgmt .
compose:
	 docker-compose -f build/docker-compose.yml up -d
remove:
	 docker container stop filemgmtservice mongoservice
	yes |  docker-compose -f build/docker-compose.yml rm
	 docker rmi -f filemgmt
status:
	docker container ls -a 
	#############################################
	docker images
	#############################################
	docker-compose -f ./build/docker-compose.yml ps 

	 