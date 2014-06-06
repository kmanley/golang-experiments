package main

import "fmt"
import "bytes"
import "encoding/gob"

func encode() *bytes.Buffer {
	indata := []interface{}{1, 2, "a", "b", "c"}
	buff := bytes.NewBufferString("")
	enc := gob.NewEncoder(buff)
	enc.Encode(indata)
	return buff
}

func decode(e *bytes.Buffer) {
	dec := gob.NewDecoder(e)
	var res interface{}
	dec.Decode(&res)
	fmt.Println("decoded: ", res)
}

func main() {
	e := encode()
	fmt.Println("encoded: ", e)
	decode(e)
}
