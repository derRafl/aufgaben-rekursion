package search

import "aufgaben-rekursion/lists"

// Find sucht in einer Liste nach dem ersten Vorkommen von x
// und gibt dessen Index zurück. Falls x nicht gefunden wird,
// wird -1 zurückgegeben.
func Find(list []int, x int) int {
	result := 0

	if lists.Empty(list) {
		return -1

	}

	if list[0] == x {
		return result
	}

	if list[0] != x {
		i := Find(list[1:], x)

		result = result + i + 1

		if i == -1 {
			result = -1
		}
		return result
	}

	return -2
}
