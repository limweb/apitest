FROM golang:1.16-alpine3.13 as install
WORKDIR /app
COPY go.mod  /app/go.mod
COPY go.sum  /app/go.sum
ENV  GIN_MODE=release
RUN  go mod download

FROM install  as builder
WORKDIR /app
COPY db  /app/db
COPY docs  /app/docs
COPY handler  /app/handler
COPY model  /app/model
COPY main.go  /app/main
COPY api.md  /app/api.md
ENV  GIN_MODE=release
RUN  CGO_ENABLED=0 GOOS=linux go build  -o main  

FROM alpine:3.13 as production
WORKDIR /app
COPY --from=builder /app/main /app/main
CMD ["./main"]

