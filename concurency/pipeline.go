/*
конвейер
  - обработка одного за раз
*/
package main

import (
	"fmt"
)

func generator(done chan interface{}, integers ...int) <-chan int {

	result := make(chan int)
	go func() {
		defer close(result)
		for _, j := range integers {
			select {
			case <-done:
				return
			case result <- j:
			}
		}
	}()
	return result
}

func multiply(done chan interface{}, n <-chan int, multiplier int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for j := range n {
			select {
			case <-done:
				return
			case result <- j * multiplier:
			}
		}
	}()

	return result
}

func add(done chan interface{}, n <-chan int, add int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for j := range n {
			select {
			case <-done:
				return
			case result <- j + add:
			}
		}
	}()

	return result
}

func main() {
	n := []int{1, 2, 3, 4}
	done := make(chan interface{})

	initChan := generator(done, n...)

	pipeline := multiply(done, add(done, multiply(done, initChan, 2), 1), 2)
	for j := range pipeline {
		fmt.Println(j)
	}

}
