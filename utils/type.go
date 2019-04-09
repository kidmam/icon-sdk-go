package utils

func isStr(value interface{}) bool {
	_, ok := value.(string)
	return ok
}

func isInteger(value interface{}) bool {
	_, ok := value.(int)
	return ok
}

func isHash(value interface{}) bool {
	_, ok := value.(string)
	return ok
}

func assertionInteger64(value interface{}) int64 {
	val, _ := value.(int)
	return int64(val)
}

func assertionString(value interface{}) string {
	val, _ := value.(string)
	return string(val)
}
