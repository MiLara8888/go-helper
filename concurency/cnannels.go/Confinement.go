package cnannelsgo

import "fmt"

func confinement() {
	// десь мы создаем экземпляр канала в лексической области действия функции chanOwner . Это
	// ограничивает область записи канала результатов замыканием, определенным под ним. Другими словами,
	// он включает аспект записи этого канала, чтобы другие горутины не могли писать в него.
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}
	// Здесь мы получаем копию int- канала, доступную только для чтения. Заявляя, что единственное
	// использование, которое нам требуется, — это доступ для чтения, мы ограничиваем использование канала
	// внутри функции потребления только чтением.
	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}
	// Здесь мы получаем аспект чтения канала и можем передать его потребителю, который ничего не может
	// делать, кроме как читать из него. И снова это ограничивает основную горутину представлением канала
	// только для чтения.
	results := chanOwner()
	consumer(results)
}
