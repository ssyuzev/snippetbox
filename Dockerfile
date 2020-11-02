FROM golang:1.15

RUN mkdir /app
COPY . /app
WORKDIR /app

# RUN go get -d -v ./...
# RUN go install -v ./...
RUN go mod download
