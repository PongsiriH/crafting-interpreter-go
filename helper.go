package main

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func isAlpha(c byte) bool {
  return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || c == '_'
}

func isAlphaNumeric(c byte) bool {
  return isAlpha(c) || isDigit(c)
}
