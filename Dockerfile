# Use an official Node.js 20 image as the base image
FROM node:20-alpine

# Set environment variables for Go installation
ENV GOLANG_VERSION 1.23.0
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 905a297f19ead44780548933e0ff1a1b86e8327bb459e92f9c0012569f76f5e3

# Install dependencies and download Go
RUN apk add --no-cache \
    build-base \
    gcc \
    musl-dev \
    libc-dev \
    git \
    curl \
    bash \
    make \
    yarn \
    && curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
    && echo "$GOLANG_DOWNLOAD_SHA256 golang.tar.gz" | sha256sum -c - \
    && tar -C /usr/local -xzf golang.tar.gz \
    && rm golang.tar.gz

# Set up Go environment variables
ENV PATH $PATH:/usr/local/go/bin

# Verify installations
RUN node --version \
    && go version

# Set a working directory
WORKDIR /app

COPY . .

RUN yarn

RUN make build

RUN chmod +x ./bin/app_prod

ENTRYPOINT ["/app/bin/main"]


