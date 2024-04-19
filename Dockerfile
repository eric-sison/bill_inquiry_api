FROM golang:1.22-alpine

WORKDIR /app

COPY . .
RUN go mod download

COPY *.go ./

RUN go build -o main main.go

EXPOSE 3276

CMD ["/bill-inquiry"]