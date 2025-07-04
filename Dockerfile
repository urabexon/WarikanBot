FROM golang:1.24

WORKDIR /app

ENV GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 5272

CMD ["go", "run", "main.go"]