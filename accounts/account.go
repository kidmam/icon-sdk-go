package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Create an account
	key, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println(err)
	}

	// Get the address
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	fmt.Println("address : " + address)

	ks := keystore.NewKeyStore(
		os.Args[1],
		keystore.LightScryptN,
		keystore.LightScryptP)
	account, err := ks.NewAccount(os.Args[2])

	if err != nil {
		log.Fatal("error decrypting key")
	}

	fmt.Printf("Account Address Hex: %s\n", "hx"+account.Address.Hex()[2:])
	fmt.Printf("Account Address String: %s\n", account.Address.String())
}
