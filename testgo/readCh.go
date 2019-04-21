package main

import (
	"fmt"
	"time"
)

func writeToChannel(c chan int, x int) {
	fmt.Println("1", x)
	c <- x
	close(c)
	fmt.Println("2", x)
}

func f1(c chan int, x int) {
	fmt.Println(x)
	c <- x
}

func f2(c chan<- int, x int) {
	fmt.Println(x)
	c <- x
}

func main() {
	c := make(chan int)
	go writeToChannel(c, 10)
	time.Sleep(1 * time.Second)
	fmt.Println("Read:", <-c)
	time.Sleep(1 * time.Second)

	_, ok := <-c
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")

	}

	go f1(c, 10)
	go f2(c, 20)
}
