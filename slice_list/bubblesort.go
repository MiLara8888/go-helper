package slicelist

// сортировка пузырьком

func (l *List) Bublsort() {
	for i := 0; i < len(l.List); i++ {

		for j := 0; j <= i; j++ {
			if l.List[j] > l.List[i] {
				l.List[j], l.List[i] = l.List[i], l.List[j]
			}
		}
	}
}
