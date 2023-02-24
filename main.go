package main

import (
	"fmt"

	"github.com/Open-Material/open-material/crypto"
)

func main() {
	wallet := crypto.NewWallet()
	//fmt.Println(*wallet)
	//fmt.Println(*wallet)

	fmt.Println("Private Key: ", string(wallet.PrivateKey))
	//fmt.Println("D: ", wallet.PrivateKey.D)
	//fmt.Println("X: ", wallet.PrivateKey.X)
	//fmt.Println("Y: ", wallet.PrivateKey.Y)
	//fmt.Println("Public Key: ", wallet.PrivateKey.PublicKey)

	//fmt.Println(crypto.KeyStoreDirExists())

	if !crypto.KeyStoreDirExists() {
		crypto.CreateKeyStoreDir()
	}
	crypto.CreateKeyStoreFile(wallet)
	walletFromKeyStore, _ := crypto.ReadFromKeyStoreFile()
	fmt.Println(*walletFromKeyStore)
}
