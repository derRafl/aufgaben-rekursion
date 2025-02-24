package lists

// Liefert das Produkt aller Elemente in list.
// Wenn list leer ist, wird 1 zurückgegeben.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func Product(list []int) int {

	result := 1

	if Empty(list) {
		return result
	}

	return result * list[0] * Product(list[1:])
}
