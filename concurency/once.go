package concurency

import (
	"fmt"
	"sync"
)

// синтаксис
func SyntOnce() {
	var count int
	increment := func() { count++ }
	decrement := func() { count-- }
	var once sync.Once
	once.Do(increment)
	once.Do(decrement)
	fmt.Printf("Count: %d\n", count)
}
