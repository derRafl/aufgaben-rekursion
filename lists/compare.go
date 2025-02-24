package lists

// Liefert true, falls die beiden Listen gleich sind.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func ListsEqual(list1, list2 []int) bool {
	// Basisfall: Beide Listen sind leer
	if Empty(list1) && Empty(list2) {
		return true
	}

	// Basisfall: Eine Liste ist leer, die andere nicht
	if Empty(list1) || Empty(list2) {
		return false
	}

	// Vergleiche erste Elemente
	if list1[0] != list2[0] {
		return false
	}

	// Rekursiver Aufruf mit Rest der Listen
	return ListsEqual(list1[1:], list2[1:])
}
