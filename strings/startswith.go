package strings

// StartsWith liefert true, falls der string s mit der Sequenz seq beginnt.
func StartsWith(s, seq string) bool {
	if s == seq {
		return true
	}

	if len(s) < len(seq) {
		return false
	}

	x := s[:len(seq)]

	if x == seq {
		return true
	}

	return false
}
