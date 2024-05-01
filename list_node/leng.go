package listnode

func (l *LinkedList) Lengh()int{
	var cou int
	node := l.Head
	for node.next!=nil {
		cou+=1
		node = node.next
	}
	return cou
}