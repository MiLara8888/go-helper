package slicelist


func merge(x []int, y []int) []int {
	a, b := 0, 0
	res := []int{}

	for {
		if len(x) == a {
			res = append(res, y[b:]...)
			break
		}
		if len(y) == b {
			res = append(res, x[a:]...)
			break
		}
		if x[a] < y[b] {
			res = append(res, x[a])
			a += 1
		} else {
			res = append(res, y[b])
			b += 1
		}
	}
	return res
}

func  MergeSort(m []int) []int {
	if len(m) <= 1 {
		return m
	}
	mid := len(m) / 2
	left := m[:mid]
	right := m[mid:]

	left = MergeSort(left)
	right = MergeSort(right)

	return merge(left, right)

}
