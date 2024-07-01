FROM golang:1.19.2-bullseye as builder

WORKDIR /server

COPY go.mod .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /serve run.go

FROM alpine:3
WORKDIR /
COPY --from=builder /serve /


ENTRYPOINT [ "/serve" ]
