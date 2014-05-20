package main

import (
	"bytes"
	"encoding/json"
	_ "fmt"
	"net/http"
	"testing"
)

type Message struct {
	Message string `json:"message"`
}

const helloWorldString = "Hello, World!"

func BenchmarkTest1(b *testing.B) {
	buff := bytes.NewBufferString("")
	for i := 0; i < b.N; i++ {
		buff.WriteString("hello, world")
	}
}

const S = "hello, world"

func BenchmarkTest2(b *testing.B) {
	buff := bytes.NewBufferString("")
	for i := 0; i < b.N; i++ {
		buff.WriteString(S)
	}
}

func BenchmarkTestJson1(b *testing.B) {
	//fmt.Println("running with b.N=", b.N)
	var data = make([]byte, 100)
	buff := bytes.NewBuffer(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buff.Reset()
		json.NewEncoder(buff).Encode(&Message{helloWorldString})
	}
	//fmt.Println(buff.String())
}

func BenchmarkTestJson2(b *testing.B) {
	//fmt.Println("running with b.N=", b.N)
	var data = make([]byte, 100)
	buff := bytes.NewBuffer(data)
	encoder := json.NewEncoder(buff)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buff.Reset()
		encoder.Encode(&Message{helloWorldString})
	}
	//fmt.Println(buff.String())
}

// 270 ns/op
func BenchmarkTestHeader1(b *testing.B) {
	h := &http.Header{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Set("Content-Type", "application/json")
	}
}

// 210 ns/op
func BenchmarkTestHeader2(b *testing.B) {
	var h http.Header = make(http.Header)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h["Content-Type"] = []string{"application/json"}
	}
}

const key = "Content-Type"

var value = []string{"application/json"}

// 94 ns/op
func BenchmarkTestHeader3(b *testing.B) {
	//var h http.Header = make(http.Header)
	h := make(http.Header)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h[key] = value
	}
}
