package main

import (
	"context"
	"os"

	. "github.com/theplant/htmlgo"
)

func main() {
	comp := Div(
		Text("123<h1>"),
	)
	Fprint(os.Stdout, comp, context.TODO())
}
