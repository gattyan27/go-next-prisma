FROM golang:1.22

RUN apt-get update -y
RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY go.mod ./

RUN go mod tidy

RUN go get github.com/prisma/prisma-client-go@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

# airを起動コマンドとして設定
CMD ["air"]