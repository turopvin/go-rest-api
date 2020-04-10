FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -v ./cmd/apiserver

EXPOSE 8080

ENTRYPOINT ["./apiserver"]
