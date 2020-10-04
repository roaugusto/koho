FROM golang:1.15

# golang specific variables
ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

ENV APP_HOME /app

WORKDIR /app

COPY . /app

RUN go mod download

RUN go build -o hokobalance .

# Command to run
ENTRYPOINT /app/hokobalance

# expose the port to run the application on
EXPOSE 3333
