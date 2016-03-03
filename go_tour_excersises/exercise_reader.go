package main

import (
	"golang.org/x/tour/reader"
	"fmt"
)

type MyReader struct{}

func (mr MyReader) Read(b []byte) (int, error){
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
