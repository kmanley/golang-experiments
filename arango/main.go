package main

import (
	"context"
	"log"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/davecgh/go-spew/spew"
	"github.com/gcmurphy/getpass"
)

type SearchTerms []string

type Attachment struct {
	Name        string
	UUID        string
	Key         string
	ContentType string
	Size        int
	Public      bool
}
type Attachments []Attachment

type Contact struct {
	Name     string
	Address  string
	Address2 string
	City     string
	State    string
	Zip      string
	Phone    string
	Email    string
	Website  string
}

type Regulator struct {
	_key         string
	LastUpdate   int // time.Time
	Name         string
	Abbreviation string
	Main         Contact
	Complaint    Contact
	RateBoard    Contact
	Notes        string
	Attachments  Attachments
	SearchTerms  SearchTerms
	NumSuppliers int // readonly TODO: is there a way to tell gorethink to not store but to also include in result of merge? `gorethink:"-"` doesn't work.
}

func main() {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
	})
	if err != nil {
		log.Fatal(err)
	}

	pwd, _ := getpass.GetPass()
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", pwd),
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	db, err := client.Database(ctx, "_system")
	if err != nil {
		log.Fatal(err)
	}

	//type Foo struct {
	//	x int // time.Time
	//}
	query := "FOR r IN regulators LIMIT 3 RETURN r"
	//query := `return {"x":1507900413446}`
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		log.Fatal(err)
	}
	//doc := &Foo{x: -2} //Regulator
	var doc Regulator
	defer cursor.Close()
	for {
		_, err := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			log.Print("no more docs")
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			spew.Dump(doc)

			log.Printf("*****************************************************************")
			log.Printf("*****************************************************************")
		}
	}

}
