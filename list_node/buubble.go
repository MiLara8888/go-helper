package listnode


func (l *LinkedList) BubbleSort(c int) {
	for i := 0; i < c-1; i++ {
		swapped := false
		first := l.Head
		for j := 0; j < c-i-1 && first.next != nil; j++ {
			if first.num > first.next.num {
				first.num, first.next.num = first.next.num, first.num
				swapped = true
			}
			first = first.next
		}
		if swapped == false {
			break
		}
	}
}
