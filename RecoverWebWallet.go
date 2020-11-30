package main

import (
	"encoding/hex"
	"fmt"
	"log"

	// "github.com/btcsuite/btcutil/hdkeychain"

	"github.com/decred/dcrd/chaincfg/v3"
	"github.com/decred/dcrd/hdkeychain/v3"
)

func main() {
	var xPrivKey string
	fmt.Print("Enter web wallet xPrivKey: ")
	fmt.Scanln(&xPrivKey)
	net := chaincfg.MainNetParams()
	key, err := hdkeychain.NewKeyFromString(xPrivKey, net)

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

	privKey, err := key.SerializedPrivKey()
	fmt.Printf("Hex seed: %s\n", hex.EncodeToString(privKey))
}
