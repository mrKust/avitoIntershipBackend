FROM golang:1.19-buster

RUN go version

ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o avitoIntershipBackend ./cmd/main/main.go

CMD ["./avitoIntershipBackend"]