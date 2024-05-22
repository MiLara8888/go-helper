package cnannelsgo

import (
	"fmt"
	"time"
)

// Здесь мы передаем канал Done функции doWork .
func doWork(done <-chan interface{}, strings <-chan string) <-chan interface{} {
	terminated := make(chan interface{})
	go func() {
		defer fmt.Println("doWork exited.")
		defer close(terminated)
		// В этой строке мы видим повсеместно используемый шаблон for-select. Один из наших
		// операторов случая проверяет, был ли сигнализирован наш канал готовности . Если да,
		// то возвращаемся из горутины.
		for {
			select {
			case s := <-strings:
				// Do something interesting
				fmt.Println(s)
				// В этой строке мы видим повсеместно используемый шаблон for-select. Один из наших
				// операторов случая проверяет, был ли сигнализирован наш канал готовности . Если да,
				// то возвращаемся из горутины.
			case <-done:
				return
			}
		}
	}()
	return terminated
}

func doNe() {

	done := make(chan interface{})
	terminated := doWork(done, nil)

	// Здесь мы создаем еще одну горутину, которая отменит горутину, созданную в doWork,
	// если пройдет более одной секунды.
	go func() {
		// Cancel the operation after 1 second.
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()
	<-terminated
	// Здесь мы соединяем горутину, порожденную из doWork , с основной горутиной.
	fmt.Println("Done.")
}
