package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	. "github.com/karrick/gorill"
)

func main() {
	r := &LineTerminatedReader{R: bytes.NewReader([]byte("123\n456"))}
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	if got, want := len(buf), 8; got != want {
		fmt.Fprintf(os.Stderr, "GOT: %v; WANT: %v\n", got, want)
		os.Exit(1)
	}
	fmt.Printf("%q\n", buf[len(buf)-1])
	// Output: '\n'
}
