package main

import (
	"fmt"
	"strconv"

	"github.com/kidmam/icon-sdk-go/iconsdk"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	iconsdk.Setendpoint("https://int-test-ctz.solidwallet.io")

	// return int64
	height, _ := iconsdk.GetLastBlockHeight()
	fmt.Println("Last block height : " + strconv.FormatInt(height, 10))

	result, err := iconsdk.GetBlock(0)

	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println("error !!")
	}

	zerolog.TimeFieldFormat = ""

	log.Print("hello world")
}
