package select_t




// Как видите, за тысячу итераций оператор select примерно в половине случаев читает из c1, и
// примерно в половине случаев — из c2.
// func chan_read() {
// 	c1 := make(chan interface{})
// 	close(c1)
// 	c2 := make(chan interface{})
// 	close(c2)
// 	var c1Count, c2Count int
// 	for i := 1000; i >= 0; i-- {
// 		select {
// 		case <-c1:
// 			c1Count++
// 		case <-c2:
// 			c2Count++
// 		}
// 	}
// 	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
// }


// func time_out() {
// 	var c <-chan int
// 	select {
// 	case <-c:
// 	case <-time.After(1 * time.Second):
// 		fmt.Println("Timed out.")
// 	}
// }





// func loop_t() {
// 	// Это позволяет горутине выполнять работу, ожидая,
// 	// пока другая горутина сообщит о результате.
// 	done := make(chan interface{})
// 	go func() {
// 		time.Sleep(5 * time.Second)
// 		close(done)
// 	}()
// 	workCounter := 0
// loop:
// 	for {
// 		select {
// 		case <-done:
// 			break loop
// 		default:
// 		}
// 		// Simulate work
// 		workCounter++
// 		time.Sleep(1 * time.Second)
// 	}
// 	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
// }
