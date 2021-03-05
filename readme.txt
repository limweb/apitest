//feature
- swagger api docs  ok
- docker support    ok
- flag arement      ok
- gorm database     ok
- jwt 
- ...

//build cammand
swag int --md ./
go run main.go  or  go build

//development use
gin -a 8000 -p 8080 run main.go 


//build docker
docker build -t <docker image tag> -f ./docker/Dockerfile .
docker run -d -p 8080:8080 <docker image tag>

//docker-compose up -d
docker-compose -f ./docker/docker-compose.yml up -d --build