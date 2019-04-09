package iconsdk

import (
	"testing"
)

func TestGetLastBlockHeight(t *testing.T) {

	Setendpoint("https://int-test-ctz.solidwallet.io")

	// GetBlock test
	latestBlock, _ := GetBlock("latest")

	genesisBlock, _ := GetBlock(0)

	if latestBlock == "" || genesisBlock == "" {
		t.Fatal("Error !")
	}

	// GetLastBlock test
	height, _ := GetLastBlockHeight()

	if height < 0 {
		t.Fatal("Height must be bigger than and equal 0.")
	}

}
