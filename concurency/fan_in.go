package main

import (
	"sync"
)

func fanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {

	var wg sync.WaitGroup
	//результирующий канал
	multiplexedStream := make(chan interface{})

	//вычитывание из канала в общий канал
	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	//разветсление
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	// ждёт пока все чтения завершаться
	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()
	return multiplexedStream
}

func main() {

}
