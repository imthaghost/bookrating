# Golang version - Docker Hub (hub.docker.com)
FROM golang:alpine3.15

# Enviornment variables
ENV APP_NAME bookrating

# Working directory
WORKDIR /go/src/${APP_NAME}

COPY . /go/src/${APP_NAME}

# Install dependecies from mod file
RUN go mod download

RUN apk add git

# Build application
RUN go build -o ${APP_NAME} ./cmd/bookrating

# Run application
CMD ./${APP_NAME}