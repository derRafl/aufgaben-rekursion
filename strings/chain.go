package strings

// Chain erwartet einen String und hängt ihn n mal hintereinander.
// Liefert das Ergebnis.
func Chain(s string, n int) string {
	result := ""
	if n > 0 {
		result = result + s
		n--
		result = result + Chain(s, n)
	}
	return result
}
