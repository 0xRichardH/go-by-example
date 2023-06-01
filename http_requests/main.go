package main

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
)

const demoURL = "https://jsonplaceholder.typicode.com"

// {
//   "userId": 1,
//   "id": 1,
//   "title": "delectus aut autem",
//   "completed": false
// }

type todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	var t todo
	ctx := context.Background()
	err := requests.
		URL(demoURL).
		Path("/todos/1").
		ToJSON(&t).
		Fetch(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t)
}
