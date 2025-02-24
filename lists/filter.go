package lists

// Liefert eine Liste mit allen Elementen aus list, die kleiner oder gleich key sind.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func FilterLess(list []int, key int) []int {
	// Gehen Sie ähnlich wie bei Remove vor:
	// Wenn die Liste leer ist, ist das Ergebnis die leere Liste.
	// Wenn das erste Element größer als key ist, ist das Ergebnis die gefilterte Restliste.
	// Wenn das erste Element kleiner oder gleich key ist, ist das Ergebnis das erste Element
	// plus die gefilterte Restliste.

	var result = []int{}

	if Empty(list) {
		return list
	}

	if list[0] < key {
		result = append(result, list[0])
		return append(result, FilterLess(list[1:], key)...)
	}

	if list[0] > key {
		return FilterLess(list[1:], key)
	}

	// TODO
	return result
}

// Liefert eine Liste mit allen Elementen aus list, die echt größer als key sind.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func FilterGreater(list []int, key int) []int {
	var result = []int{}

	if Empty(list) {
		return list
	}

	if list[0] > key {
		result = append(result, list[0])
		return append(result, FilterGreater(list[1:], key)...)
	}

	if list[0] < key {
		return FilterGreater(list[1:], key)
	}

	// TODO
	return result
}
