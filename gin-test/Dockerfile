FROM golang:1.24.4-bullseye

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod tidy
RUN go mod download

COPY . .

RUN go build -o bin/server main.go

ENTRYPOINT [ "./bin/server" ]