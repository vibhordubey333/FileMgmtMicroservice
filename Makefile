all: docker compose

docker:
	sudo docker build -f build/Dockerfile -t filemgmt .
compose:
	sudo docker-compose -f build/docker-compose.yml up -d
remove:
	sudo docker container stop filemgmtservice mongoservice
	yes | sudo docker-compose -f build/docker-compose.yml rm
	sudo docker rmi -f filemgmt
status:
	docker container ls -a 
	#############################################
	docker images
	#############################################
	docker-compose -f ./build/docker-compose.yml ps 

	 