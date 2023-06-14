## Redis Pub/Sub

- Start redis instance

```
docker-compose up
```

- Start Subscriber

```
go run ./cmd/sub/main.go
```

- Start publishing message

```
go run ./cmd/pub/main.go
```
