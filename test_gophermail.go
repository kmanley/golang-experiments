package main

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/gcmurphy/getpass"
	"github.com/kmanley/gophermail"
)

func main() {
	m := &gophermail.Message{}
	m.SetFrom("Utiliz <kevin@utiliz.co>")
	//m.AddTo("Kevin Manley <kevin.manley@gmail.com>")
	m.AddTo("Tom Place <tom@utiliz.co>")
	m.Subject = "Test multi-part message with atch sent from Golang via sendgrid"
	m.HTMLBody = `
<html>
  <body>
    <h2>Test</h2>
	Did it work?
	<hr/>
	<table border="1">
	<tr>
	  <th>Col1</th><th>Col2</th>
	</tr>
	<tr>
	  <td>Foo</td><td>Bar</td>
	</tr>
	</table>
	<br/>
	<img src="https://cdn.scratch.mit.edu/static/site/projects/thumbnails/1095/8698.png"/>
  </body>
</html>
`
	m.Body = "This is the plain text body for those without an HTML capable email client"
	//fmt.Println(m.Bytes())
	atchname := "logo-resistor-400x400.png"
	data, _ := os.Open(atchname)
	atch := &gophermail.Attachment{Name: atchname,
		Data: data}
	m.Attachments = append(m.Attachments, *atch)

	pwd, _ := getpass.GetPass()

	a := smtp.PlainAuth("", "utiliz",
		pwd, "smtp.sendgrid.net")

	err := gophermail.SendMail("smtp.sendgrid.net:587", a, m)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("ok")
	}

}
