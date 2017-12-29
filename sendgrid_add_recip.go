package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
)

func main() {
	host := "https://api.sendgrid.com"
	apiKey := os.Getenv("sendgrid_temp_key")
	if apiKey == "" {
		fmt.Println("no sendgrid_temp_key var found")
		os.Exit(1)
	}
	request := sendgrid.GetRequest(apiKey, "/v3/contactdb/recipients", host)
	request.Method = "POST"
	request.Body = []byte(` [
		  {
		    "email": "kevin.manley@gmail.com",
		    "first_name": "Kevin",
		    "last_name": "Manley"
		  }
		]`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Printf("error: %s\n", err)
	} else {
		log.Println("SUCCESS!")
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
