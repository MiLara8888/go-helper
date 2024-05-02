package main

import (
	"fmt"

	"sync"
	"time"
)

func listen(name string, data map[string]string, c *sync.Cond) {
	c.L.Lock()
	c.Wait()

	fmt.Printf("[%s] %s\n", name, data["key"])

	c.L.Unlock()
}

func broadcast(name string, data map[string]string, c *sync.Cond) {
	time.Sleep(time.Second)

	c.L.Lock()

	data["key"] = "value"

	fmt.Printf("[%s] данные получены\n", name)

	c.Signal()
	c.L.Unlock()
}

func main() {
	data := map[string]string{}

	cond := sync.NewCond(&sync.Mutex{})


	go listen("слушатель 1", data, cond)
	go listen("слушатель 2", data, cond)

	go broadcast("источник", data, cond)



}