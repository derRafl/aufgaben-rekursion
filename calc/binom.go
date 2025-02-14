package calc

// BinomialCoefficient erwartet zwei Zahlen n und k und liefert den
// Binomialkoeffizienten "n über k".
func BinomialCoefficient(n, k int) int {
	if n < 1 || k < 1 || (n == 1 && k == 1) || n == k {
		return 1
	}

	return BinomialCoefficient(n-1, k) + BinomialCoefficient(n-1, k-1)
}

/**
	func BinomialCoefficient(n, k int) int {

		res := (fak(n)) / (fak(k) * fak(n-k))

		return res
	}

	func fak(n int) int {
		if n == 0 {
			return 1
		}

		return (n * fak(n-1))
	}
**/
