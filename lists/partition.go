package lists

// Liefert zwei Listen:
// - Eine, die alle Elemente aus list enthält, die kleiner oder gleich key sind.
// - Eine, die alle übrigen Elemente aus list enthält.
func Partition(list []int, key int) ([]int, []int) {
	// Verwende Kopien von list, damit die ursprüngliche Liste nicht verändert wird.
	var l1 = []int{}
	var l2 = []int{}

	if Empty(list) {
		return l1, l2
	}

	if list[0] <= key {
		l1 = append(l1, list[0])
	}

	if list[0] > key {
		l2 = append(l2, list[0])
	}

	l11, l22 := Partition(list[1:], key)
	return append(l1, l11...), append(l2, l22...)

}
