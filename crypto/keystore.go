package crypto

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// KeyStoreDirExists checks if the keystore directory exists.
func KeyStoreDirExists() bool {
	_, err := os.Stat("keystore")
	return !os.IsNotExist(err)
}

// CreateKeyStoreDir creates the keystore directory.
func CreateKeyStoreDir() error {
	err := os.Mkdir("keystore", 0700)

	if err != nil {
		return err
	}
	return nil
}

// CreateKeyStoreFile creates the keystore file, and writes the private key and address to it.
// TODO: Add password encryption.
func CreateKeyStoreFile(wallet *Wallet) error {
	f, err := os.Create("keystore/keystore")

	if err != nil {
		return err
	}

	defer f.Close()

	keystore, err := json.Marshal(wallet)

	if err != nil {
		return err
	}

	_, err = f.WriteString(string(keystore))

	if err != nil {
		return err
	}

	return nil
}

// ReadFromKeyStoreFile reads the keystore file and returns the private key and address (Wallet struct).
// TODO: make this function return the private key and address separately.
func ReadFromKeyStoreFile() (*Wallet, error) {
	body, err := ioutil.ReadFile("keystore/keystore")

	if err != nil {
		return nil, err
	}

	wallet := &Wallet{}
	err = json.Unmarshal(body, wallet)
	//fmt.Println(wallet)

	if err != nil {
		return nil, err
	}

	return wallet, nil
}
