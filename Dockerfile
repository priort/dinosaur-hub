FROM golang:1.16

COPY . .

EXPOSE 8084

RUN go build dinosaur-info.go

CMD ["./dinosaur-info"]