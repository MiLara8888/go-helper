package slicelist

// сортировка вставками
func (l *List) InsertionSort() {
	for i := 1; i < len(l.List); i++ {
		for j := i; j >= 1 && l.List[j] < l.List[j-1]; j-- {
			l.List[j], l.List[j-1] = l.List[j-1], l.List[j]
		}
	}
}
