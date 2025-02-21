package internal

func isDigit(v byte) bool {
	return '0' <= v && v <= '9'
}

func isAlpha(v byte) bool {
	return ('a' <= v && v <= 'z') || ('A' <= v && v <= 'Z')
}

func isAlphanumeric(v byte) bool {
	return isAlpha(v) || isDigit(v)
}

func isEqual(a, b any) bool {
	return a == b
}

func isTruthy(val any) bool {
	return (val != nil) && (val != false)
}
