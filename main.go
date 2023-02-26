package main

import (
	"fmt"
	"math/big"

	"github.com/Open-Material/open-material/crypto"
	"github.com/Open-Material/open-material/proposal"
)

func main() {
	//wallet := crypto.NewWallet()
	//fmt.Println(*wallet)
	//fmt.Println(*wallet)

	//fmt.Println("Private Key: ", string(wallet.PrivateKey))
	//fmt.Println("D: ", wallet.PrivateKey.D)
	//fmt.Println("X: ", wallet.PrivateKey.X)
	//fmt.Println("Y: ", wallet.PrivateKey.Y)
	//fmt.Println("Public Key: ", wallet.PrivateKey.PublicKey)

	//fmt.Println(crypto.KeyStoreDirExists())

	if !crypto.KeyStoreDirExists() {
		crypto.CreateKeyStoreDir()
	}
	//crypto.CreateKeyStoreFile(wallet)
	walletFromKeyStore, _ := crypto.ReadFromKeyStoreFile()
	fmt.Println(*walletFromKeyStore)
	proposalWithoutSig := proposal.Proposal{
		Nonce:        1,
		From:         &walletFromKeyStore.Address,
		Title:        "Test",
		Pos:          new(big.Int).SetInt64(1),
		Height:       new(big.Int).SetInt64(12),
		FWHMLeft:     new(big.Int).SetInt64(35),
		DSpacing:     new(big.Int).SetInt64(19),
		RelIntensity: new(big.Int).SetInt64(100),
	}

	fmt.Println("Hash of the proposal", proposalWithoutSig.Hash())
	parsedPriv, err := crypto.ToECDSA(walletFromKeyStore.PrivateKey)
	if err != nil {
		fmt.Println("Error parsing private key")
	}

	signedProposal, err := proposal.SignProposal(&proposalWithoutSig, parsedPriv)
	if err != nil {
		fmt.Println("Error signing proposal")
	}

	fmt.Printf("%+v\n", *signedProposal)
}
