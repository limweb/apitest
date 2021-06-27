#feature
- swagger api docs  ok
- docker support    ok
- flag arement      ok
- gorm database     ok
- jwt               ok
- vuetable          ok  
- hasOne            ok
- hsaMany           ok
- Gorm Relationship ok

#***************************************************
# SNIPPET ใช้   Ext:  servit-snippet v0.0.15 ขึ้นไป //
#***************************************************

# Clone
##git clone https://github.com/limweb/apitest
##cd apitest
##go mod download



#https://localhost:8080/api/istabs/vuetable?page=2&per_page=4&sort=typecode|asc,volue|desc&kw=xxx&filter=field|aaaa,field2|bbbbbb


#ก่อนจะทำอะไรให้ติดตั้ง  2  ตัว 1  ไว้ทำ hot reload   2. ไว้ gen apidoc
#เป็นตัว swag ไว้ gen apidoc
go get -u github.com/swaggo/swag/cmd/swag

#gin ไว้ทำ hot reload
#https://medium.com/thipwriteblog/golang-live-reload-4e1f96648b80  วิธีใช้
go get github.com/codegangsta/gin

#build cammand
swag int --md ./
go run main.go  or  go build

#development use
gin -a 8000 -p 8080 run main.go 

#build docker
docker build -t <docker image tag> -f ./docker/Dockerfile .
docker run -d -p 8080:8080 <docker image tag>

#docker-compose up -d
docker-compose -f ./docker/docker-compose.yml up -d --build