FROM  golang:1.17.1-stretch

WORKDIR /app

COPY go.mod . 
COPY go.sum .
RUN go mod download

COPY . .

RUN `go mod tidy \
  && go build src/main.go`

CMD ["./main"]
