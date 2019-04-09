package keycreate

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"os"
)

const (
	veryLightScryptN = 2
	veryLightScryptP = 1
)

func main() {
	keyjson, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Print(err)
	}
	password := ""
	address := common.HexToAddress("")
	fmt.Print(address)
	// Do a few rounds of decryption and encryption
	for i := 0; i < 3; i++ {
		// Try a bad password first
		if _, err := keystore.DecryptKey(keyjson, password+"bad"); err == nil {
			fmt.Printf("test %d: json key decrypted with bad password", i)
		}
		// Decrypt with the correct password
		key, err := keystore.DecryptKey(keyjson, password)
		if err != nil {
			fmt.Printf("test %d: json key failed to decrypt: %v", i, err)
		}
		if key.Address != address {
			fmt.Printf("test %d: key address mismatch: have %x, want %x", i, key.Address, address)
		}
		// Recrypt with a new password and start over
		password += "new data appended"
		if keyjson, err = keystore.EncryptKey(key, password, veryLightScryptN, veryLightScryptP); err != nil {
			fmt.Printf("test %d: failed to recrypt key %v", i, err)
		}
	}

}
