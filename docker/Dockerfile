FROM golang:1.16-alpine3.13 as install
WORKDIR /app
COPY go.mod  /app/go.mod
COPY go.sum  /app/go.sum
ENV  GIN_MODE=release
RUN  go mod download

FROM install  as builder
COPY ./main.go  /app/main.go
COPY ./api.md   /app/api.md
COPY ./db       /app/db
COPY ./utils   /app/utils
COPY ./webserver   /app/webserver
COPY ./controllers   /app/controllers
COPY ./docs     /app/docs
COPY ./middleware   /app/middleware
COPY ./models    /app/models
COPY ./publics   /app/publics
COPY ./services   /app/services
WORKDIR /app
ENV  GIN_MODE=release
RUN  CGO_ENABLED=1 GOOS=linux go build  -o main  

# FROM alpine:3.13 as production
# WORKDIR /app
# COPY --from=builder /app /app
CMD ["./main"]
