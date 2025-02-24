package strings

// ContainsChain liefert true, falls s eine Kette von count aufeinanderfolgenden
// Vorkommen von symbol enthält.
func ContainsChain(s, symbol string, count int) bool {
	// TODO
	counter := 0

	if s == symbol {
		return true
	}

	if len(s) <= len(symbol) {
		return count == 0
	}

	x := s[:len(symbol)]

	if x == symbol {
		counter++
	}

	if Contains(s[1:], symbol) {
		counter++
	}

	return counter >= count
}
