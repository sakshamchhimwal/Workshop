FROM golang:1.23.0-alpine3.20 AS server
WORKDIR /home/workshop/server
COPY . /home/workshop/server/
RUN go mod download
RUN go build -o ./build/server
CMD ["./build/server"]