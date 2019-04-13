package main

import (
	"fmt"
	"os"
	"reflect"
)

type a struct {
	X int
	Y float64
	Z string
}

type b struct {
	F int
	G int
	H string
	I float64
}

func main() {
	x := 100
	xRepl := reflect.ValueOf(&x).Elem()
	xType := xRepl.Type()
	fmt.Printf("The type of x %s\n", xType)

	A := a{100, 200.12, "Struct a"}
	B := b{1, 2, "Struct b", -1.2}

	var r reflect.Value

	argumetns := os.Args

	if len(argumetns) == 1 {
		r = reflect.ValueOf(&A).Elem()
	} else {
		r = reflect.ValueOf(&B).Elem()
	}

	iType := r.Type()
	fmt.Printf("i Type: %s\n", iType)
	fmt.Printf("The %d fields of %s are:\n", r.NumField(), iType)

	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("Filed name: %s ", iType.Field(i).Name)
		fmt.Printf("with type: %s ", r.Field(i).Type())
		fmt.Printf("and value: %v\n", r.Field(i).Interface())
	}
}
