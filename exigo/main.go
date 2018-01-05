package main

import (
	"flag"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/glog"
	exigo "github.com/kmanley/go-exigo-client"
	exigoapi "github.com/kmanley/go-exigo-client/api"
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

func testCreateCustomer(client *exigo.Client) {
	req := &exigoapi.CreateCustomerRequest{
		CustomerType: 1,
		FirstName:    "Johnny",
		LastName:     "Smithson",
		Email:        "john@aol.com",
		Phone:        "203-555-1212",
		Notes:        "great guy!",
		MainAddress1: "100 Nowhere Lane",
		MainCity:     "Tucson",
		MainState:    "AZ",
		MainZip:      "85990",
		MainCountry:  "US",
		//LoginName:    "johnny",
	}
	res, err := client.CreateCustomer(req)
	if err != nil {
		glog.Errorf("CreateCustomer error: %s", err)
	} else {
		glog.Info("success")
		spew.Dump(res)
	}
}

func testCreateConsultant(client *exigo.Client) {
	/*
		req := &exigoapi.CreateCustomerRequest{
			CustomerType:       3,
			FirstName:          "Michael",
			MiddleName:         "Lee",
			LastName:           "Martin",
			Email:              "mike@aol.com",
			Phone:              "888-555-1212",
			Notes:              "A real jerk!",
			MainAddress1:       "333 Oak St",
			MainCity:           "Manchester",
			MainState:          "CT",
			MainZip:            "06880",
			MainCountry:        "US",
			SponsorID:          5,
			InsertUnilevelTree: true,
			//LoginName:    "johnny",
		}
	*/
	req := &exigoapi.CreateCustomerRequest{
		CustomerType:       3,
		FirstName:          "Debbie",
		MiddleName:         "",
		LastName:           "Downline",
		Email:              "dd@aol.com",
		Phone:              "888-355-1212",
		MainAddress1:       "444 Elm St",
		MainCity:           "Vernon",
		MainState:          "CT",
		MainZip:            "06890",
		MainCountry:        "US",
		SponsorID:          20,
		InsertUnilevelTree: true,
		LoginName:          "deb",
		LoginPassword:      "debpwd",
		CanLogin:           true,
	}

	res, err := client.CreateCustomer(req)
	if err != nil {
		glog.Errorf("CreateCustomer error: %s", err)
	} else {
		glog.Info("success")
		spew.Dump(res)
	}
}

func testUpdateCustomer(client *exigo.Client) {
	req := &exigoapi.UpdateCustomerRequest{CustomerType: 3}
	err := client.UpdateCustomer(20, req)
	if err != nil {
		glog.Errorf("UpdateCustomer error: %s", err)
	} else {
		glog.Info("success")
	}
}

func main() {
	//DT1 := time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)
	//DT2 := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	//url := "https://v5live4.exigo.com/admin6/fb6632a52597447d8a986e4dde101ca4/"
	flag.Parse()
	client := exigo.NewClient("", "kevinm", os.Getenv("kpass"), "vivnetwork")
	//testGetCustomer(client)
	//testCreateCustomer(client)
	testCreateConsultant(client)
	//testUpdateCustomer(client)
}
