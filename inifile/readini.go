package main

import (
	"fmt"
	"github.com/robfig/config"
)


func main() {
c, _ := config.ReadDefault("example.ini")

res1,_ := c.String("service-1", "url")
// result is string "http://www.example.com/some/path"
fmt.Println(res1)

res2,_ := c.Int("service-1", "maxclients")
// result is int 200
fmt.Println(res2)

res3, _ := c.Bool("service-1", "delegation")
// result is bool true
fmt.Println(res3)

res4, _ := c.String("service-1", "comments")
// result is string "This is a multi-line\nentry"
fmt.Println(res4)

}

