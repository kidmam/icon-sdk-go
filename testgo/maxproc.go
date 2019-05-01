package main

import (
	"fmt"
	"runtime"
)

func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}

func main() {
	fmt.Println("GOMAXPROC: %d\n", getGOMAXPROCS())
}
