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
