package database

import (
	"fmt"

	badger "github.com/dgraph-io/badger/v3"
)

type Database struct {
	Db *badger.DB
}

func (d *Database) View(key []byte) {
	//opts := badger.DefaultOptions("/tmp/badger")
	//opts = opts.WithLogger(nil)
	//db, err := badger.Open(opts)

	err := d.Db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)

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

func (d *Database) Write(key []byte, value []byte) {
	//opts := badger.DefaultOptions("/tmp/badger")
	//opts = opts.WithLogger(nil)
	//db, err := badger.Open(opts)
	txn := d.Db.NewTransaction(true)
	err := txn.SetEntry(badger.NewEntry(key, value))

	if err != nil {
		panic(err)
	}

	err = txn.Commit()

	if err != nil {
		panic(err)
	}
}
