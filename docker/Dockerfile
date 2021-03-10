FROM golang:latest as install
WORKDIR /app
COPY go.mod  /app/go.mod
COPY go.sum  /app/go.sum
ENV  GIN_MODE=release
RUN  go mod download

FROM install  as builder
COPY ./main.go  /app/main.go
COPY ./api.md   /app/api.md
COPY ./webapi.db  /app/webapi.db
COPY ./config.yml /app/config.yml
COPY ./db       /app/db
COPY ./config /app/config
COPY ./utils   /app/utils
COPY ./webserver   /app/webserver
COPY ./controllers   /app/controllers
COPY ./docs     /app/docs
COPY ./middleware   /app/middleware
COPY ./models    /app/models
COPY ./publics   /app/publics
COPY ./services   /app/services
ENV  GIN_MODE=release
WORKDIR /app
RUN  GOOS=linux go build  -o main  
# RUN  CGO_ENABLED=0 GOOS=linux go build  -o main  
FROM golang:latest as production
WORKDIR /app
COPY --from=builder /app/main /app/main
COPY ./publics  /app/publics
COPY ./utils/private.key  /app/utils/private.key
COPY ./utils/public.key  /app/utils/public.key
COPY ./config.yml  /app/config.yml
CMD ["./main"]
