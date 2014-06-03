package main

import "fmt"
import "bytes"
import "encoding/json"

func encode() {
	indata := []interface{}{1, 2, "a", "b", "c"}
	buff := bytes.NewBufferString("")
	enc := json.NewEncoder(buff)
	enc.Encode(indata)
	fmt.Println(buff.String())
}

func decode() {
	data := bytes.NewBufferString(`[1,2,"a","b","c"]`)
	dec := json.NewDecoder(data)
	var res interface{}
	dec.Decode(&res)
	fmt.Println(res)

}

func main() {
	encode()
	decode()
}
