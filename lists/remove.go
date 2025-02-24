package lists

// Liefert eine Liste, die alle Elemente aus list enthält,
// außer dem an Stelle pos.
// Wenn pos außerhalb des Bereichs der Liste liegt, wird die
// ursprüngliche Liste zurückgegeben.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func RemoveElement(list []int, pos int) []int {
	var empty = []int{}
	var result = []int{}

	if Empty(list) {
		return empty
	}

	if pos == 0 {

	}

	if pos < 0 {
		return append(result, list...)
	}

	if pos > 0 {
		result = append(result, list[0])
	}
	return append(result, RemoveElement(list[1:], pos-1)...)

}
