package main

import (
	"fmt"

	"github.com/enriquebris/goconcurrentqueue"
)

type AnyStruct struct {
	Field1 string
	Field2 int
}

func main() {
	queue := goconcurrentqueue.NewFIFO()

	queue.Enqueue("any string value")
	queue.Enqueue(5)
	queue.Enqueue(AnyStruct{Field1: "hello world", Field2: 15})

	// will output: 3
	fmt.Printf("queue's length: %v\n", queue.GetLen())

	item, err := queue.Dequeue()
	if err != nil {
		fmt.Println(err)
		return
	}

	// will output "any string value"
	fmt.Printf("dequeued item: %v\n", item)

	// will output: 2
	fmt.Printf("queue's length: %v\n", queue.GetLen())

}
