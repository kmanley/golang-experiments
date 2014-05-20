package main

import (
	"os"
	"fmt"
	"gosqlite.googlecode.com/hg/sqlite"
)

type Article struct {
	Id    int
	Title string
	Body  string
	Date  string
}

func main() {
	fmt.Println("\n********* Reading from a SQLite3 Database **********")
	db := "test.db"

	conn, err := sqlite.Open(db)
	if err != nil {
		fmt.Println("Unable to open the database: %s", err)
		os.Exit(1)
	}

	defer conn.Close()

	conn.Exec("CREATE TABLE articles(id INTEGER PRIMARY KEY AUTOINCREMENT, title VARCHAR(200), body TEXT, date TEXT);")

	insertSql := `INSERT INTO articles(title, body, date) VALUES("This is a Test Article Title.",
        "MO, dates are a pain.  I spent considerable time trying to decide how best to
        store dates in my app(s), and eventually chose to use Unix times (integers).
        It seemed an easy choice as I program in Perl and JavaScript.",
        "12/05/2010");`

	err = conn.Exec(insertSql)
	if err != nil {
		fmt.Println("Error while Inserting: %s", err)
	}

	selectStmt, err := conn.Prepare("SELECT id, title, body, date FROM articles;")
	err = selectStmt.Exec()
	if err != nil {
		fmt.Println("Error while Selecting: %s", err)
	}

	if selectStmt.Next() {
		var article Article

		err = selectStmt.Scan(&article.Id, &article.Title, &article.Body, &article.Date)
		if err != nil {
			fmt.Printf("Error while getting row data: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Id => %d\n", article.Id)
		fmt.Printf("Title => %s\n", article.Title)
		fmt.Printf("Body => %s\n", article.Body)
		fmt.Printf("Date => %s\n", article.Date)
	}
}
