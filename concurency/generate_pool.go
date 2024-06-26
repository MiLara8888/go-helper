/*
конвейер
  - обработка одного за раз
*/
package main

import "fmt"

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

// Этот этап конвейера удалит только первое число элементов из входящего потока значений.
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

func main() {
	done := make(chan interface{})
	defer close(done)
	for num := range take(done, repeat(done, 1), 10) {
		fmt.Printf("%v ", num)
	}
}
