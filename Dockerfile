FROM golang:1.23

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

COPY /data/lazymeal/prd/app_db /data/lazymeal/prd/app_db

RUN make sitemap && mv sitemap.xml ./public/assets/sitemap.xml

RUN go build -ldflags="-s -w" -v -o /usr/local/bin/app_prod ./cmd/app/main.go

ENTRYPOINT ["app_prod"]
