FROM golang:1.19.2-bullseye

WORKDIR /server

COPY . .

RUN go mod download

RUN go build -o /goauthdocker

EXPOSE 8089

CMD ["/goauthdocker"]