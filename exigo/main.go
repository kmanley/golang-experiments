package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/gcmurphy/getpass"
	"github.com/kmanley/golang-experiments/exigo/myservice"
)

func main() {
	//url := "https://v5live4.exigo.com/admin6/fb6632a52597447d8a986e4dde101ca4/"
	url := "http://api.exigo.com/3.0/"
	username := "kevinm"
	password, _ := getpass.GetPass()
	auth := &myservice.BasicAuth{username, password}
	client := myservice.NewExigoApiSoap(url, true, auth)
	client.AddHeader(&myservice.ApiAuthentication{LoginName: username, Password: password, Company: "vivnetwork"})
	req := &myservice.GetCustomersRequest{CustomerID: 4}
	res, err := client.GetCustomers(req)
	if err != nil {
		fmt.Println("error: %s", err)
		return
	}
	fmt.Println("success!")
	spew.Dump(res)
}
