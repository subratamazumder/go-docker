# Go Docker
## Description
Sample REST API using Golang HTTP module that stores data into DynamoDB
## Test In Local
```
go run eprescription-reg-service.go
curl -is http://0.0.0.0:8081/ep-registration-service/health
curl -X POST -is http://0.0.0.0:8081/ep-registration-service/registrations \
   -H "Content-Type: application/json" \
   -d '{"firstName": "cyano", "lastName": "mazumder"}'
 ~/workspace/go-docker   master  go run eprescription-reg-service.go
2021/10/03 02:00:48 HTTP Go Server is Listening on  192.168.1.101 : 8081
2021/10/03 02:00:54 Request received from 127.0.0.1:59138
2021/10/03 02:01:03 Request received from 127.0.0.1:59147

```
## Build Executable
```
GO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o eprescription-reg-service
```
## Build Docker Image
Make sure to build image in this format <hub-user>/<repo-name>[:<tag>]
```
docker build -t dockersubrata/eprescription-reg-service-image:1.0 .
Sending build context to Docker daemon    7.7MB
Step 1/5 : FROM alpine:latest
 ---> 14119a10abf4
Step 2/5 : WORKDIR service
 ---> Using cache
 ---> 7c9b158123f9
Step 3/5 : COPY eprescription-reg-service /service/
 ---> Using cache
 ---> 75964a78194b
Step 4/5 : EXPOSE 8081
 ---> Using cache
 ---> d62e760b2c9a
Step 5/5 : CMD ["/service/eprescription-reg-service"]
 ---> Using cache
 ---> 883af1c83603
Successfully built 883af1c83603

```
## Test Docker Image
```
 192  ~/workspace/aws-learning/eks   master ●  docker run dockersubrata/eprescription-reg-service-image:1.0
2021/10/03 02:38:33 HTTP Go Server is Listening on  c5edb02c8e9d : 8081
```
## Publish To Docker Hub
```
192  ~/workspace/go-docker   master ●  docker login
Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
Username: dockersubrata    
Password: 
Login Succeeded
 192  ~/workspace/go-docker   master ●  docker push dockersubrata/eprescription-reg-service-image:1.0
The push refers to repository [docker.io/dockersubrata/eprescription-reg-service-image]
6a5e13e9e27d: Pushed 
e89aa6f400d2: Pushed 
e2eb06d8af82: Mounted from library/alpine 
1.0: digest: sha256:eda4bcb3e0c8b1fdb84f341d639e25b3fb17203c32fa2c3ab67870961d75f1ea size: 946
 192  ~/workspace/go-docker   master ●  
```
## Tips
To get faster use `./publish-new-image.sh <image-version>`
