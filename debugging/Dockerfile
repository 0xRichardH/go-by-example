FROM golang:1.20.4

RUN go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

CMD ["dlv", "trace", "--ebpf", "foo"]
