package main

import (
	"fmt"

	"github.com/moonrhythm/passwordtool"
)

func main() {
	hashed, err := passwordtool.Hash("superman")
	if err != nil {
		// ...
	}
	fmt.Println(hashed)

	err = passwordtool.Compare(hashed, "superman")
	if err == passwordtool.ErrMismatched {
		// not equal
	}
	if err != nil {
		// ...
	}
}
