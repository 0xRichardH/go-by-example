package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/rueidis"
)

func main() {
	ctx := context.Background()

	client := createRedisClient()
	defer client.Close()

	// client.B().Subscribe().Channel("ch1", "ch2").Build()
	err := client.Receive(ctx, subscribeCommand(client), msgHandlerFn)
	if err != nil {
		log.Fatal(err)
		return
	}
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

func subscribeCommand(client rueidis.Client) rueidis.Completed {
	return client.B().Subscribe().Channel("ch1", "ch2").Build()
}

func msgHandlerFn(msg rueidis.PubSubMessage) {
	fmt.Printf("Message received: %v\n", msg)
}
