package main

import (
	"fmt"

	"github.com/h2non/filetype"
)

var fooType = filetype.NewType("foo", "foo/foo")

func fooMatcher(buf []byte) bool {
	return len(buf) > 1 && buf[0] == 0x01 && buf[1] == 0x02
}

func main() {
	// Register the new matcher and its type
	filetype.AddMatcher(fooType, fooMatcher)

	// Check if the new type is supported by extension
	if filetype.IsSupported("foo") {
		fmt.Println("New supported type: foo")
	}

	// Check if the new type is supported by MIME
	if filetype.IsMIMESupported("foo/foo") {
		fmt.Println("New supported MIME type: foo/foo")
	}

	// Try to match the file
	fooFile := []byte{0x01, 0x02}
	kind, _ := filetype.Match(fooFile)
	if kind == filetype.Unknown {
		fmt.Println("Unknown file type")
	} else {
		fmt.Printf("File type matched: %s\n", kind.Extension)
	}
}
