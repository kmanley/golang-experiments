package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"os"
)

func numKeys(bucket *bolt.Bucket) int {
	stats := bucket.Stats()
	return stats.KeyN
}

func main() {

	filename := "boltdb_stats_issue.db"
	os.Remove(filename)
	db, err := bolt.Open(filename, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_ = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("test"))
		if err != nil {
			return err
		}

		// no keys yet, expect 0
		fmt.Println("expected 0, got ", numKeys(bucket))

		bucket.Put([]byte("key1"), []byte("hello"))
		// just added 1 key, expect stats to return 1 but it returns 0
		fmt.Println("expected 1, got ", numKeys(bucket))

		return nil
	})

	_ = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("test"))
		// update tx finished, stats now returns 1 for number of keys
		fmt.Println("expected 1, got ", numKeys(bucket))
		return nil
	})

}
