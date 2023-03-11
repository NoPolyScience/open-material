package main

import (
	"fmt"
	"math/big"

	"github.com/Open-Material/open-material/crypto"
	"github.com/Open-Material/open-material/database"
	"github.com/Open-Material/open-material/proposal"
)

func main() {
	localdb := database.Database{}
	localdb.View()
	localdb.Write()
	localdb.View()
}

func createNewWallet() {
	wallet := crypto.NewWallet()
	fmt.Println(*wallet)

	//fmt.Println(crypto.KeyStoreDirExists())

	if !crypto.KeyStoreDirExists() {
		crypto.CreateKeyStoreDir()
	}
	crypto.CreateKeyStoreFile(wallet)
}

func signProposal() {
	walletFromKeyStore, _ := crypto.ReadFromKeyStoreFile()
	fmt.Println(*walletFromKeyStore)

	proposalWithoutSig := proposal.Proposal{
		Nonce:        1,
		Proposer:     &walletFromKeyStore.Address,
		Title:        "Test",
		Pos:          new(big.Int).SetInt64(1),
		Height:       new(big.Int).SetInt64(12),
		FWHMLeft:     new(big.Int).SetInt64(35),
		DSpacing:     new(big.Int).SetInt64(19),
		RelIntensity: new(big.Int).SetInt64(100),
	}

	//fmt.Println("Hash of the proposal", proposalWithoutSig.Hash())
	parsedPriv, err := crypto.ToECDSA(walletFromKeyStore.PrivateKey)
	if err != nil {
		fmt.Println("Error parsing private key")
	}

	signedProposal, err := proposal.SignProposal(&proposalWithoutSig, parsedPriv)
	if err != nil {
		fmt.Println("Error signing proposal")
	}

	fmt.Printf("%+v\n", *signedProposal)

	fmt.Println(proposal.ValidateProposal(signedProposal))
}
