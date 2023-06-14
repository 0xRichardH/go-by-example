package main

import (
	"context"
	"fmt"

	"github.com/redis/rueidis"
)

func main() {
	ctx := context.Background()
	client := createRedisClient()
	defer client.Close()

	client.Do(ctx, client.B().Publish().Channel("ch1").Message("hello, world").Build())
	client.Do(ctx, client.B().Publish().Channel("ch2").Message("hey hey").Build())
	client.Do(ctx, client.B().Publish().Channel("ch2").Message("hey hey").Build())
	fmt.Println("Done.")
}

func createRedisClient() rueidis.Client {
	client, err := rueidis.NewClient(
		rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}},
	)
	if err != nil {
		panic(err)
	}
	return client
}
