# Use an official Node.js 20 image as the base image
FROM node:20 AS NODE_BUILDER

RUN node --version

COPY package.json .

RUN npm install

COPY . .

RUN npx tailwindcss -i app/assets/app.css -o ./public/assets/styles.css
RUN npx esbuild app/assets/index.js --bundle --outdir=public/assets


FROM golang:1.23

RUN go version

WORKDIR /app

COPY --from NODE_BUILDER . .

RUN go build -ldflags="-s -w" -o bin/app_prod cmd/app/main.go

RUN chmod +x bin/app_prod

ENTRYPOINT ["/app/bin/main"]


