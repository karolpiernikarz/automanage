package cache

import (
	"log"

	"github.com/dgraph-io/badger/v4"
)

var mdb *badger.DB

func Init() {
	// open in memory database
	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	// check if error
	if err != nil {
		log.Fatal(err)
	}
	// set mdb
	mdb = db
}
