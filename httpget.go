package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main1() {
	for {
		url := "localhost:8080/foo"
		fmt.Println("getting", url)
		http.Get(url)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	client := &http.Client{}
	ctr := 0
	for {
		url := "http://localhost:8080/foo"
		fmt.Println(ctr, "getting", url)
		resp, err := client.Get(url)
		if err != nil {
			fmt.Println(err)
		} else {
			bytes, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(bytes))
			resp.Body.Close()
		}
		time.Sleep(10 * time.Second)
		ctr += 1
	}
}
