FROM golang:1.19-alpine

WORKDIR /app

COPY . /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/gila-test.bin /app/interface/http/main.go 

EXPOSE 8080

CMD [ "/app/gila-test.bin", "0.0.0.0" ]