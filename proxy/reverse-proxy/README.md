## Reverse Proxy

> https://dev.to/b0r/implement-reverse-proxy-in-gogolang-2cp4

- To run and test the origin server

```
go run ./cmd/origin/main.go
```

```
curl -i localhost:3000/hey
```

- To run and test the reverse proxy server

```
go run ./cmd/proxy/main.go
```

```
curl -i localhost:8080/hey
```
