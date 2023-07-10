FROM golang:1.19

WORKDIR /app

COPY . .

RUN go mod download

RUN go build . 

EXPOSE 8081

CMD ["./main"]