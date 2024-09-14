FROM golang:1.23

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN echo "DB_NAME=app_db" > .env
RUN echo "DB_DRIVER=sqlite3" >> .env
RUN echo "DB_USER=" >> .env
RUN echo "DB_HOST=" >> .env
RUN echo "MIGRATION_DIR=app/db/migrations" >> .env

RUN make sitemap && cp sitemap.xml ./public/assets/sitemap.xml

RUN go build -ldflags="-s -w" -v -o /usr/local/bin/app_prod ./cmd/app/main.go

ENTRYPOINT ["app_prod"]
