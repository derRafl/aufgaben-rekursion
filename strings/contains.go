package strings

// Contains prüft, ob der String s die Sequenz seq enthält.
func Contains(s, seq string) bool {

	if len(s) < len(seq) {
		return false
	}

	x := s[:len(seq)]

	if x == seq {
		return true
	}

	return Contains(s[1:], seq)

}
