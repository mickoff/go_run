package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	fields = strings.Fields(s)
	for word := range fields {
		m[word] = 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
