package main

import (
	"fmt"
	"github.com/davidbanham/human_duration"
	"time"
)

func main() {
	example := time.Hour*24 + time.Minute*4 + time.Second*8

	fmt.Println(human_duration.String(example, "second")) // 1 day 4 minutes 8 seconds
	fmt.Println(human_duration.String(example, "minute")) // 1 day 4 minutes
	fmt.Println(human_duration.String(example, "day"))    // 1 day

	day := time.Hour * 24
	year := day * 365

	longExample := year*4 + day*2

	fmt.Println(human_duration.String(longExample, "second")) // 4 years 2 days
}
