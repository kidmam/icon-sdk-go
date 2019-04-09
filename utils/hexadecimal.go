package utils

func is0xPrefixed(value string) bool {
	if isStr(value) && len(value) > 2 {
		return value[0:2] == "0x"
	}
	return false
}

func Add0xPrefix(value string) string {
	if is0xPrefixed(value) {
		return value
	}
	return "0x" + value
}
