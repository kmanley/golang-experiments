package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/robertkrimen/otto"
)

func main() {

	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("usage: run <jscript>")
	}

	jspath := args[0]
	jscript, err := ioutil.ReadFile(jspath)
	if err != nil {
		log.Fatalf("failed to load js file %s: %s", jspath, err)
	}

	vm := otto.New()
	_, err = vm.Run(string(jscript))
	if err != nil {
		log.Fatalf("script execution error: %s", err)
	}
}
