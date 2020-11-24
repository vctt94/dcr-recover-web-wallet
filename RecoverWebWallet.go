package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/btcsuite/btcutil/hdkeychain"
)

func main() {
	var xPrivKey string
	fmt.Print("Enter web wallet xPrivKey: ")
	fmt.Scanln(&xPrivKey)
	key, err := hdkeychain.NewKeyFromString(xPrivKey)

	if err != nil {
		log.Fatal(err)
		return
	}

	if key.Depth() != 0 {
		log.Fatal("This xPrivKey is not a seed")
		return
	}
	if key.ParentFingerprint() != 0 {
		log.Fatal("This xPrivKey is not a seed")
		return
	}

	privKey, err := key.ECPrivKey()
	fmt.Printf("Hex seed: %s\n", hex.EncodeToString(privKey.Serialize()))
}
