package lex

func IsNumber(r rune) bool {
	switch {
	case r >= '0' && r <= '9':
		return true
	default:
		return false
	}
}

// isSpace returns true if the rune is a tab or space.
func IsSpace(r rune) bool {
	return r == '\u0009' || r == '\u0020'
}

// isNameBegin returns true if the rune is an alphabet or an '_' or '~'.
func IsNameBegin(r rune) bool {
	switch {
	case r >= 'a' && r <= 'z':
		return true
	case r >= 'A' && r <= 'Z':
		return true
	case r == '_':
		return true
	//case r == '~':
	//	return true
	default:
		return false
	}
}
