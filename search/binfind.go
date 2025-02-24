package search

// FindSorted sucht in einer sortierten Liste nach dem ersten Vorkommen von x.
// Falls x nicht gefunden wird, wird -1 zurückgegeben.
// Da die Liste sortiert ist, wird die binäre Suche verwendet.
func FindSorted(list []int, x int) int {
	if len(list) == 0 {
		return -1
	}

	l := len(list) / 2

	if list[l] == x {
		return l
	}

	if list[l] < x {
		// Rekursiver Aufruf mit rechter Hälfte
		result := FindSorted(list[l+1:], x)
		if result == -1 {
			return -1
		}
		return l + 1 + result
	}

	// Rekursiver Aufruf mit linker Hälfte
	return FindSorted(list[:l], x)
}
