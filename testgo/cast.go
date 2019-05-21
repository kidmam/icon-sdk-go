package main

import "github.com/spf13/cast"

func main() {
	cast.ToString("mayonegg")         // "mayonegg"
	cast.ToString(8)                  // "8"
	cast.ToString(8.31)               // "8.31"
	cast.ToString([]byte("one time")) // "one time"
	cast.ToString(nil)                // ""

	var foo interface{} = "one more time"
	cast.ToString(foo) // "one more time"
}
