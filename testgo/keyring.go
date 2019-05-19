package main

import (
	"fmt"
	"github.com/99designs/keyring"
)

func main() {
	ring, _ := keyring.Open(keyring.Config{
		ServiceName: "example",
	})

	_ = ring.Set(keyring.Item{
		Key:  "foo",
		Data: []byte("secret-bar"),
	})

	i, _ := ring.Get("foo")

	fmt.Printf("%s", i.Data)
}
