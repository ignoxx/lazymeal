FROM golang:1.23

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./

RUN go mod download && go mod verify

RUN go install github.com/a-h/templ/cmd/templ@latest

COPY . .

RUN templ generate

RUN cd public && cp *.png assets/

RUN go build -ldflags="-s -w" -v -o /usr/local/bin/app_prod ./cmd/app/main.go

ENTRYPOINT ["app_prod"]
