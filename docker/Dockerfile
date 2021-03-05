FROM golang:1.16-alpine3.13 as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY db .
COPY handler .
COPY docs .
COPY model .
COPY api.md .
RUN  go mod download
RUN  CGO_ENABLED=0 GOOS=linux go build -o main ./...

FROM alpine:3.13 as production

WORKDIR /app
COPY --from=builder /app .
CMD ["main"]