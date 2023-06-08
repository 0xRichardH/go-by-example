FROM almalinux:8.5

RUN dnf install epel-release dnf-plugins-core -y
RUN dnf --enablerepo=powertools install clang go make libbpf-devel -y

RUN go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

CMD ["~/go/bin/dlv", "trace", "--ebpf", "foo"]
