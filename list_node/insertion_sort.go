package listnode


func (l *LinkedList) InsertionSort() {

	headPre := &Node{next: l.Head}

	cur := l.Head
	for cur != nil && cur.next != nil {
		p := cur.next
		if cur.num <= p.num {
			cur = p
			continue
		}
		cur.next = p.next
		pre, next := headPre, headPre.next

		for next.num < p.num {
			pre = next
			next = next.next
		}
		pre.next = p
		p.next = next

	}
	l.Head = headPre.next.next
}


