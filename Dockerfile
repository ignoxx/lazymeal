FROM golang:1.23

WORKDIR /app

COPY . .

RUN go version

RUN go build -ldflags="-s -w" -o bin/app_prod cmd/app/main.go

RUN chmod +x bin/app_prod

ENTRYPOINT ["/app/bin/main"]


