FROM golang:1.19.2-bullseye

WORKDIR /server

COPY . .

RUN go mod download

EXPOSE 8088

CMD go run ./run.go
