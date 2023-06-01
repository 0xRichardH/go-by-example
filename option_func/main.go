package main

import "fmt"

func main() {
	u := CreateUser(
		WithUserName("Alice"),
		WithUserAge(29),
	)
	fmt.Printf("%+v\n", u)
}

type User struct {
	name string
	age  int
}

type CreateUserOptionFn func(u *User)

func CreateUser(opts ...CreateUserOptionFn) *User {
	u := &User{}
	for _, opt := range opts {
		opt(u)
	}
	return u
}

func WithUserName(name string) CreateUserOptionFn {
	return func(u *User) {
		u.name = name
	}
}

func WithUserAge(age int) CreateUserOptionFn {
	return func(u *User) {
		u.age = age
	}
}
