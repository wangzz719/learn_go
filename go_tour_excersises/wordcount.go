package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	var fields = strings.Fields(s)
	var result = make(map[string]int)
	for _, value := range fields{
		result[value] +=1
	}
	return result
}

func main() {
	fmt.Println(WordCount("Word Count Word Count"))
}
