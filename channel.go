package main

import (
	"fmt"
	"sync"
)

func main() {
	orderCh := make(chan int)
	accountCh := make(chan int)
	doneCh := make(chan bool)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer func() {
			wg.Done()
			close(orderCh)
		}()
		orderCh <- 1
	}()

	go func() {
		defer func() {
			wg.Done()
			close(accountCh)
		}()
		accountCh <- 2
	}()

	go func() {
		wg.Wait()
		close(doneCh)
	}()

	for {
		select {
		case order, ok := <-orderCh:
			if ok {
				fmt.Println(order)
			}
		case account, ok := <-accountCh:
			if ok {
				fmt.Println(account)
			}
		case <-doneCh:
			return
		}
	}
}
