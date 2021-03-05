//build cammand
swag int --md ./
go run main.go

//development use
gin -a 8000 -p 8080 run main.go 


//build docker
docker build -t apitest -f ./docker/Dockerfile .
docker run -d -p 8080:8080 apitest


//docker-compose up -d
docker-compose -f ./docker/docker-compose.yml up -d 