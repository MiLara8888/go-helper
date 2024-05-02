/*
две горутины ждут (cond.Wait), пока третья запишет данные в общий словарь и сообщит об этом с помощью cond.Broadcast.
*/
package concurency
import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func Listen(name string, data map[string]string, c *sync.Cond) {
	c.L.Lock()
	c.Wait()
	fmt.Printf("[%s] %s\n", name, data["key"])
	c.L.Unlock()
}

func Broadcast(name string, data map[string]string, c *sync.Cond) {
	time.Sleep(time.Second)
	c.L.Lock()
	data["key"] = "value"
	fmt.Printf("[%s] данные получены\n", name)
	c.Broadcast()
	c.L.Unlock()
}

func Work() {
	data := map[string]string{}
	cond := sync.NewCond(&sync.Mutex{})
	go Listen("слушатель 1", data, cond)
	go Listen("слушатель 2", data, cond)
	go Broadcast("источник", data, cond)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

// словари передаются по ссылке, и слушатели смогут прочитать данные, которые будут записаны в переданный словарь
// чтобы программа не завершилась до того, как горутины закончат работу, мы ждём сигнал os.Interrupt (ctrl + c в терминале)
// sync.NewCond принимает интерфейс sync.Locker (мьютекс удовлетворяет ему) и возвращает ссылку на инстанс sync.Cond.
// Внутри broadcast мы вызываем time.Sleep(time.Second), чтобы гарантировать, что cond.Broadcast будет вызван после вызова cond.Wait внутри listen
// Вызов cond.Wait обязательно нужно вызывать до вызова cond.Broadcast, иначе слушатели повиснут навсегда

//-------------------------------------------------------------------------------------------------------------------

// /*
// создаем приложение
// с графическим интерфейсом и кнопкой на нем. Мы хотим зарегистрировать произвольное количество функций, которые
// будут запускаться при нажатии этой кнопки. Cond идеально подходит для этого, поскольку мы можем использовать его
// метод Broadcast для уведомления всех зарегистрированных обработчиков
// */
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// // Мы определяем тип Button , который содержит условие Clicked.
// type Button struct {
// 	Clicked *sync.Cond
// }

// // Здесь мы определяем удобную функцию, которая позволит нам регистрировать функции для обработки
// // сигналов из условия. Каждый обработчик запускается в своей собственной горутине, и подписка не
// // завершится до тех пор, пока не будет подтверждено, что эта горутина запущена.
// func Subscribe(c *sync.Cond, fn func()) {
// 	var goroutineRunning sync.WaitGroup
// 	goroutineRunning.Add(1)
// 	go func() {
// 		goroutineRunning.Done()
// 		c.L.Lock()
// 		defer c.L.Unlock()
// 		c.Wait()
// 		fn()
// 	}()
// 	goroutineRunning.Wait()
// }

// func Click() {
// 	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}
// 	// Здесь мы устанавливаем обработчик поднятия кнопки мыши.
// 	// Он, в свою очередь, вызывает Broadcast при нажатии кнопки, чтобы сообщить всем обработчикам, что кнопка мыши нажата
// 	var clickRegistered sync.WaitGroup
// 	// Здесь мы создаем WaitGroup. Это делается только для того, чтобы гарантировать, что наша программа не завершит работу до
// 	// того, как произойдет запись на стандартный вывод .
// 	clickRegistered.Add(3)
// 	Subscribe(button.Clicked, func() {
// 		fmt.Println("Maximizing window.")
// 		clickRegistered.Done()
// 	})
// 	Subscribe(button.Clicked, func() {
// 		fmt.Println("Displaying annoying dialog box!")
// 		clickRegistered.Done()
// 	})
// 	Subscribe(button.Clicked, func() {
// 		fmt.Println("Mouse clicked.")
// 		clickRegistered.Done()
// 	})
// 	// Далее мы моделируем, как пользователь поднимает кнопку мыши после нажатия кнопки приложения.
// 	button.Clicked.Broadcast()
// 	clickRegistered.Wait()
// }

