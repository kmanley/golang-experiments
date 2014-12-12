package main

import (
	_ "fmt"
	"github.com/boltdb/bolt"
	"log"
)

func DoUpdate(db *bolt.DB) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return err
		}

		err = b.Put([]byte("answer"), []byte("42"))
		if err != nil {
			return err
		}
		return nil
	})
	return err

}

func DoRead(db *bolt.DB) error {
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		log.Println("b is %s", b)
		//if err != nil {
		//	return err
		//}

		v := b.Get([]byte("answer"))
		//if err != nil {
		//	return err
		//}
		log.Printf("the value is %s", v)
		return nil
	})
	return err
}

func main() {

	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("bolt-test.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//DoUpdate(db)
	DoRead(db)

}
