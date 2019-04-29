package main

import (
	"fmt"
	"time"
)

// Print text after the given time has expired.
// When done, the wait channel is closed.
func Publish(text string, delay time.Duration) (wait <-chan struct{}) {
	ch := make(chan struct{})
	go func() {
		time.Sleep(delay)
		fmt.Println("BREAKING NEWS:", text)
		close(ch) // Broadcast to all receivers.
	}()
	return ch
}

func main() {
	wait := Publish("Channels let goroutines communicate.", 5*time.Second)
	fmt.Println("Waiting for news...")
	<-wait
	fmt.Println("Time to leave.")
}
