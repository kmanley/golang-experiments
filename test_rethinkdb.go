package main

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	"log"
)

func main() {
	sess, err := r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
		AuthKey: "",
	})
	if err != nil {
		log.Fatalf("Error connecting to DB: %s", err)
	}
	curs, err := r.Db("test").Table("user").Get("foobar@gmail.com").Run(sess)
	if err != nil {
		log.Fatalf("Error:", err)
	}
	var response string
	err = curs.One(&response)
	fmt.Println(response)

}
