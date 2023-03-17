package main

import (
	"fmt"
	"math/big"

	"github.com/Open-Material/open-material/crypto"
	"github.com/Open-Material/open-material/database"
	"github.com/Open-Material/open-material/proposal"
	"github.com/Open-Material/open-material/utils"

	badger "github.com/dgraph-io/badger/v3"
)

func main() {
	//signProposal()
	//tryDatabase()
	signAddDb()
}

func signAddDb() {
	walletKeyStore, _ := crypto.ReadFromKeyStoreFile()
	proposalWithoutSig := proposal.Proposal{
		Nonce:        1,
		Proposer:     &walletKeyStore.Address,
		Title:        "Test",
		Pos:          new(big.Int).SetInt64(1),
		Height:       new(big.Int).SetInt64(12),
		FWHMLeft:     new(big.Int).SetInt64(35),
		DSpacing:     new(big.Int).SetInt64(19),
		RelIntensity: new(big.Int).SetInt64(100),
	}

	parsedPriv, err := crypto.ToECDSA(walletKeyStore.PrivateKey)
	if err != nil {
		fmt.Println("Error parsing private key")
	}

	signedProposal, err := proposal.SignProposal(&proposalWithoutSig, parsedPriv)
	if err != nil {
		fmt.Println("Error signing proposal")
	}
	rlpWithSig, _ := utils.EncodeProposalWithSig(signedProposal)

	fmt.Println(rlpWithSig)
	handleDatabase([]byte("Cu2S"), rlpWithSig)

}

func handleDatabase(name []byte, proposal []byte) {
	opts := badger.DefaultOptions("/tmp/badger")
	opts = opts.WithLogger(nil)
	db, err := badger.Open(opts)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	localdb := database.Database{Db: db}

	//localdb.View(name)
	localdb.Write(name, proposal)

	localdb.View(name)
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
	rlp, err := utils.EncodeProposal(&proposalWithoutSig)
	if err != nil {
		fmt.Println("Error encoding RLP")
	}
	fmt.Println("RLP Encoded", rlp)

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
