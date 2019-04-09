package singleton

import (
	"fmt"
	"testing"
)

func Testsingleton(m *testing.M) {
	s := New()

	s["this"] = "that"

	s2 := New()

	fmt.Println("This is ", s2["this"])
	// This is that
	m.Run()
}
