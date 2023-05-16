package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	sliceCh := converToSlice(arr)
	squareCh := square(sliceCh)

	for v := range squareCh {
		fmt.Println(v)
	}
}

func converToSlice(arr []int) <-chan int {
	outCh := make(chan int)

	go func() {
		for _, v := range arr {
			outCh <- v
		}
		close(outCh)
	}()

	return outCh
}

func square(inCh <-chan int) <-chan int {
	outCh := make(chan int)
	go func() {
		for v := range inCh {
			outCh <- v * v
		}
		close(outCh)
	}()

	return outCh
}
