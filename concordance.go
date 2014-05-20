package main

import (
  "strings"
  "sort"
  "fmt"
  "io/ioutil"
)
var print = fmt.Println

func main() {
    filename := "corpus.txt"
    bs, err := ioutil.ReadFile(filename)
    if err != nil {
        return
    }
    str := string(bs)
    words := strings.Split(str, " ")    
    conc := make(map[string]int, 100)
    for _, word := range words {
        word := strings.ToLower(word)
        print(word)
        count, exists := conc[word]
        if exists {
            conc[word] = count + 1
        } else {
            conc[word] = 1
        }
        print("-----------------------------------")
        //print(conc)
        //keys := []string
        //for key, value := range conc {
        //    keys[
        keys := make(sort.StringSlice, 0)
        for key, _ := range conc {
            //print(key, value)
            keys = append(keys, key)
        }
        //print(keys)
        sort.Sort(keys)
        //print(keys)
        for _, key := range keys {
            print(key, conc[key])
        }
        
        
    }
       
}
