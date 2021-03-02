package lexer

func isNewline(ch rune) bool {
	return ch == '\n'
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t'
}

func isAlpha(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isIdentChar(ch rune) bool {
	if ch >= 0x1f600 { // >= 😀
		return true
	}

	if isAlpha(ch) || isDigit(ch) {
		return true
	}

	for _, identch := range "!?@_$" {
		if identch == ch {
			return true
		}
	}

	return false
}

func limit(str []byte, position, length int) []byte {
	min := position - length

	if min < 0 {
		return str
	}

	return str[min : position+1]
}
