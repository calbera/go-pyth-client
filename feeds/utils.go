package feeds

// AddHexPrefix adds a '0x' prefix to a string if it doesn't already have one.
func AddHexPrefix(str string) string {
	if has0xPrefix(str) {
		return str
	}
	return "0x" + str
}

// has0xPrefix validates str begins with '0x' or '0X'.
func has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}
