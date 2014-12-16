package main

import (
	"bytes"
	"encoding/gob"
	_ "fmt"
	"github.com/boltdb/bolt"
	"log"
	"time"
)

type JobControl struct {
	MaxConcurrency         uint32
	StartTime              time.Time
	ContinueJobOnTaskError bool
	RemoteCurDir           string
	WorkerNameRegex        string
	//	CompiledWorkerNameRegex *regexp.Regexp
	// TODO: consider OSRegex as well, to limit to Workers matching a particular OS/version
	ProcessPriority int
	// TODO: later
	//AssignSingleTaskPerWorker bool
	//TaskWorkerAssignment      map[string][]uint32
	JobPriority          int8    // higher value means higher priority
	JobTimeout           float64 // seconds
	TaskTimeout          float64 // seconds
	TaskSeemsHungTimeout uint32
	AbandonedJobTimeout  uint32
	MaxTaskReallocations uint8
}

func (this *JobControl) ToBytes() ([]byte, error) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(this)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

func (this *JobControl) FromBytes(data []byte) error {
	buff := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buff)
	err := dec.Decode(this)
	if err != nil {
		return err
	}
	return nil
}

/*
func JobControlFromBytes(data []byte) (*JobControl, error) {
	buff := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buff)
	ctrl := &JobControl{}
	err := dec.Decode(ctrl)
	if err != nil {
		return nil, err
	}
	return ctrl, nil
}
*/

func DoUpdate(db *bolt.DB) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err != nil {
			return err
		}

		ctrl := &JobControl{MaxConcurrency: 316, StartTime: time.Now(), ContinueJobOnTaskError: true, RemoteCurDir: "/some/dir"}
		data, err := ctrl.ToBytes()
		if err != nil {
			log.Println("err: ", err)
			return err
		}
		err = b.Put([]byte("answer"), data)
		if err != nil {
			log.Println("err: ", err)
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
		//log.Printf("v=%s", v)
		//if err != nil {
		//	return err
		//}
		//ctrl, err := JobControlFromBytes(v)
		ctrl := &JobControl{}
		err := ctrl.FromBytes(v)
		if err != nil {
			log.Println("error", err)
			return err
		}
		log.Printf("the value is %+v", ctrl)
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

	DoUpdate(db)
	DoRead(db)

}
