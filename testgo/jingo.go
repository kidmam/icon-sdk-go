package main

//https://github.com/bet365/jingo
import (
	"fmt"
	"github.com/bet365/jingo"
)

// sample struct we'll encode
type MyPayload struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	ID   int    // anything we don't annotate doesn't get emitted.
}

// Create an encoder, letting it know which type of struct we're going to be encoding.
// You only do this once per type!
var enc = jingo.NewStructEncoder(MyPayload{})

func main() {
	// now lets encode something
	p := MyPayload{
		Name: "Mr Payload",
		Age:  33,
	}

	// pull a buffer from the pool and pass it along with the struct to Marshal
	buf := jingo.NewBufferFromPool()
	enc.Marshal(&p, buf)

	fmt.Println(buf.String()) // {"name":"Mr Payload","age":33}

	// return the buffer to the pool now we're done
	buf.ReturnToPool()
}
