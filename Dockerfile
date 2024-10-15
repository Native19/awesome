FROM golang:1.22
WORKDIR /awesome

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /my_server

EXPOSE 8080
CMD ["/my_server"]