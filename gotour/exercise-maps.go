package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	result := make(map[string]int)

	for _, v := range strings.Fields(s) {
		fmt.Printf("value[%s]\n", v)
		result[v]++
	}

	return result
}

func main() {
	wc.Test(WordCount)
}
