/*
конвейер
  - обработка одного за раз
*/
package main

import (
	"fmt"
	"math/rand"
)

// Эта функция будет бесконечно повторять переданные вами значения, пока вы не прикажете ей остановиться.
func repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}

func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func repeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

func main() {
	done := make(chan interface{})
	defer close(done)
	rand := func() interface{} { return rand.Int() }
	for num := range take(done, repeatFn(done, rand), 10) {
		fmt.Println(num)
	}
}
