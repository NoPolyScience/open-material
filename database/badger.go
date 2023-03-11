package database

import (
	"fmt"

	badger "github.com/dgraph-io/badger/v3"
)

type Database struct {
	db *badger.DB
}

func (d *Database) View() {
	opts := badger.DefaultOptions("/tmp/badger")
	opts = opts.WithLogger(nil)
	db, err := badger.Open(opts)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("key"))

		if err != nil {
			fmt.Println(err)
			return nil
		}
		val, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", string(val))
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func (d *Database) Write() {
	opts := badger.DefaultOptions("/tmp/badger")
	opts = opts.WithLogger(nil)
	db, err := badger.Open(opts)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	txn := db.NewTransaction(true)
	err = txn.SetEntry(badger.NewEntry([]byte("key"), []byte("value")))

	if err != nil {
		panic(err)
	}

	err = txn.Commit()

	if err != nil {
		panic(err)
	}
}
