package main

import (
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/glog"
	exigo "github.com/kmanley/go-exigo-client"
)

func testGetCustomer(client *exigo.Client) {
	res, err := client.GetCustomerByID(20)
	if err != nil {
		glog.Errorf("GetCustomerByID error: %s", err)
	} else {
		glog.Info("success")
		spew.Dump(res)
	}
}

/*
func testCreateCustomer(client *exigo.ExigoApiSoap) {
	req := &exigo.CreateCustomerRequest{
		CustomerType: 1,
		//FirstName: "Alfred",
		//LastName:  "Ignatius",
		Email: "alfred@aol.com",
		Phone: "203-555-1212",
		Notes: "Created via API",
	}

	res, err := client.CreateCustomer(req)
	spew.Dump(res)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	fmt.Println("success!")

}
*/

func main() {
	//DT1 := time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)
	//DT2 := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	//url := "https://v5live4.exigo.com/admin6/fb6632a52597447d8a986e4dde101ca4/"
	client := exigo.NewClient("", "kevinm", os.Getenv("kpass"), "vivnetwork")
	testGetCustomer(client)
	//testCreateCustomer(client)
}
