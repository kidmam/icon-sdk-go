package main

import (
	"errors"
	"fmt"
)

func main() {
	e1 := errors.New(fmt.Sprintf("Could not open file"))
	e2 := fmt.Errorf("Cound not open file")

	fmt.Println(fmt.Sprintf("Type of error 1: %T", e1))
	fmt.Println(fmt.Sprintf("Type of error 2: %T", e2))
}
