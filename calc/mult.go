package calc

// Liefert das Produkt von x und y.
func Mult(x, y int) int {

	if x == 0 || y == 0 {
		return 0
	}

	return x + Mult(x, y-1)
}
