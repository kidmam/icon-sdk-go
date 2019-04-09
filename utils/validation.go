package utils

func IsHexBlockHash(value interface{}) bool {
	return isStr(value) && is0xPrefixed(assertionString(value))
}

func IsBlockHeight(value interface{}) bool {
	return isInteger(value) // ...
}

func IsPredefinedBlockValue(value interface{}) bool {
	return isStr(value) && value == "latest"
}

func AssertionInteger(value interface{}) int64 {
	return assertionInteger64(value)
}

func AssertionString(value interface{}) string {
	return assertionString(value)
}
