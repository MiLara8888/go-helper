package slicelist

import (
	"fmt"
	"math/rand"
)

type List struct {
	List []int
}

func (l *List) Fill(n int) {

	for i := n; i > 0; i-- {
		l.List = append(l.List, rand.Int())
	}
}

func (n *List) Console() {
	for _, j := range n.List{
		fmt.Println(j)
	}
}
