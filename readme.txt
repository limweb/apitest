//build cammand
swag int --md ./
go run main.go

//development use
gin -a 8000 -p 8080 run main.go 