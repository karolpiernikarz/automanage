package cache

import (
	"time"

	"github.com/dgraph-io/badger/v4"
)

func DeleteKey(key []byte) (err error) {
	err = mdb.Update(func(txn *badger.Txn) error {
		err := txn.Delete(key)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return
}

func IsKeyExist(key []byte) (exist bool) {
	err := mdb.View(func(txn *badger.Txn) error {
		_, err := txn.Get(key)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false
	}
	return true
}

func GetValueFromKey(key string) (value string, err error) {
	err = mdb.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			value = string(val)
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return value, err
}

func SetKeyValue(key []byte, value []byte, ttl time.Duration) (err error) {
	err = mdb.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry(key, value).WithTTL(ttl)
		err := txn.SetEntry(e)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return
}

func GetValueFromKeysWithPrefix(prefix string) (keys map[string]string, err error) {
	err = mdb.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte(prefix)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				keys[string(k)] = string(v)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return keys, err
}

func GetKeysFromPrefix(prefix string) (keys []string, err error) {
	err = mdb.View(func(txn *badger.Txn) (err error) {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte(prefix)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			k := item.Key()
			keys = append(keys, string(k))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return
}
