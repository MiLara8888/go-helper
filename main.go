/*
создаем приложение
с графическим интерфейсом и кнопкой на нем. Мы хотим зарегистрировать произвольное количество функций, которые
будут запускаться при нажатии этой кнопки. Cond идеально подходит для этого, поскольку мы можем использовать его
метод Broadcast для уведомления всех зарегистрированных обработчиков
*/
package main

import (
	"fmt"
	"sync"
)

// Мы определяем тип Button , который содержит условие Clicked.
type Button struct {
	Clicked *sync.Cond
}

// Здесь мы определяем удобную функцию, которая позволит нам регистрировать функции для обработки
// сигналов из условия. Каждый обработчик запускается в своей собственной горутине, и подписка не
// завершится до тех пор, пока не будет подтверждено, что эта горутина запущена.
func Subscribe(c *sync.Cond, fn func()) {
	var goroutineRunning sync.WaitGroup
	goroutineRunning.Add(1)
	go func() {
		goroutineRunning.Done()
		c.L.Lock()
		defer c.L.Unlock()
		c.Wait()
		fn()
	}()
	goroutineRunning.Wait()
}

func main() {
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}
	// Здесь мы устанавливаем обработчик поднятия кнопки мыши.
	// Он, в свою очередь, вызывает Broadcast при нажатии кнопки, чтобы сообщить всем обработчикам, что кнопка мыши нажата
	var clickRegistered sync.WaitGroup
	// Здесь мы создаем WaitGroup. Это делается только для того, чтобы гарантировать, что наша программа не завершит работу до
	// того, как произойдет запись на стандартный вывод .
	clickRegistered.Add(3)
	Subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})
	Subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})
	Subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})
	// Далее мы моделируем, как пользователь поднимает кнопку мыши после нажатия кнопки приложения.
	button.Clicked.Broadcast()
	clickRegistered.Wait()
}
