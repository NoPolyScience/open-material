package crypto

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/tyler-smith/go-bip39"
)

//type Wallet struct {
//	PrivateKey ecdsa.PrivateKey
//	Address    common.Address
//}

type Wallet struct {
	PrivateKey string
	Address    common.Address
}

func NewWallet() *Wallet {
	priv, address := GenerateAddress()
	return &Wallet{PrivateKey: KeyToString(priv), Address: address}
}

// GenerateAddress generates a new private key and returns the address.
func GenerateAddress() (*ecdsa.PrivateKey, common.Address) {
	priv, _ := GeneratePrivKey()
	address := KeyToAddress(&priv.PublicKey)

	return priv, address
}

func ToECDSA(privKey string) (*ecdsa.PrivateKey, error) {
	key, err := ethcrypto.HexToECDSA(privKey)
	return key, err
}

// GeneratePrivKey generates a new private key.
func GeneratePrivKey() (*ecdsa.PrivateKey, error) {
	key, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	return key, err
}

// ShowPubKey returns the public key of a private key.
func ShowPubKey(privKey *ecdsa.PrivateKey) crypto.PublicKey {
	return privKey.Public()
}

// KeyToAddress returns the address of a public key.
func KeyToAddress(pubKey *ecdsa.PublicKey) common.Address {
	pubBytes := elliptic.Marshal(secp256k1.S256(), pubKey.X, pubKey.Y)
	return common.BytesToAddress(ethcrypto.Keccak256(pubBytes[1:])[12:])
}

// KeyToString returns the string representation of a private key.
func KeyToString(privKey *ecdsa.PrivateKey) string {
	return hex.EncodeToString(ethcrypto.FromECDSA(privKey))
}

// NewMnemonic generates a new mnemonic phrase using the BIP39 standard.
func NewMnemonic() (string, error) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return "", err
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}

	return mnemonic, nil
}
