FROM golang:1.24-alpine
WORKDIR /myapp

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download
COPY . .

CMD ["/myapp/bin/api"]
EXPOSE 8080