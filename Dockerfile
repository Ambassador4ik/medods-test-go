FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go generate ./ent

RUN go build -o app github.com/Ambassador4ik/medods-test-go/cmd/server

CMD ["./app"]