package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("hydra.db", 0600, nil)
	PrintFatalError(err)
	defer func() {
		db.Close()
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	op := flag.String("op", "", "Add or Get operation?")
	user := flag.String("u", "", "Enter username")
	pass := flag.String("p", "", "Enter password")
	flag.Parse()

	switch strings.ToUpper(*op) {
	case "ADD":
		if len(*user) != 0 && len(*pass) != 0 {
			err = addToVault(db, "PasswordVault", *user, *pass)
			PrintFatalError(err)
			fmt.Printf("Key (%s) Vaue (%s) Added ", *user, *pass)
		}
	case "GET":
		if len(*user) != 0 {
			pass, err := getFromVault(db, "PasswordVault", *user)
			PrintFatalError(err)
			fmt.Println("Password found: ", pass)
		}
	default:
	}
}

func addToVault(db *bolt.DB, bucket, key, value string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		err = b.Put([]byte(key), []byte(value))
		return err
	})
}

func getFromVault(db *bolt.DB, bucket, key string) (string, error) {
	value := ""
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		v := b.Get([]byte(key))
		value = string(v)
		return nil
	})
	return value, err
}

func PrintFatalError(err error) {
	if err != nil {
		panic(fmt.Sprint("Error happend. %s", err))
	}
}
