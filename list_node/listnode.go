package listnode

import (
	"fmt"
	"math/rand"
)

type Node struct {
	num  int
	next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Fill(n int) {

	var k = &Node{}
	for i := n; i > 0; i-- {
		if i == n {
			l.Head = &Node{}
			l.Head.num = rand.Int()
			k = &Node{}
			l.Head.next = k

		} else {
			k.num = rand.Int()
			k.next = &Node{}
			k = k.next
		}
	}
}

func (n *Node) Console() {
	for {
		if n.next == nil {
			break
		}
		fmt.Println(n.num)
		n = n.next
	}
}


