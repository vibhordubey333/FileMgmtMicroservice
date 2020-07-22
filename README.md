# FileMgmtMicroservice
File Management Microservice to save , fetch , list tar from MongoDB . Using Docker for contanerization.
# FileManagementMicroservice
## To save,retrieve,list archieve files from MongoDB this microservice is about. For contanerization Docker is used.
 
# Run.
0.Clone the project.

1.Inside project directory, fire command **"make all**".

  To remove the container,images,docker-compose simply fire command "**make remove**".

  To see the port exposed and to verify whether the container is up fire command "**make status**".

2.Swagger is present in "api" directory.

3.Testcases are present under test directory.

4."build" container files related to docker, docker-compose and k8s files. "docker-compose.yml" two services one of  
  them is for MongoDB.

5.If you want to run without docker fire command "**go run cmd/file-mgmt-microservice-server/main.go**"

6.Future scope for this microservice.

  A) Deploy using HelmChart.
  B) Add functionality to delete the archieves.
  C) Serve concurrent request.
  D) Add GUI.
 