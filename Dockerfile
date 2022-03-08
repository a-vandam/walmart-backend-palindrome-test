FROM golang:1.17

WORKDIR usr/src/cmd

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["main.go"]