FORM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY db .
COPY handler .
COPY docs .
COPY model .
COPY api.md .

RUN  go mod download

ENV PORT  8080

RUN go build

CMD ["apitest"]