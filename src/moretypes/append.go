package main

import (
	"fmt"
)

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
// reference
// https://ja.stackoverflow.com/questions/55764/go-%E3%81%AE-slice-%E3%81%AE-cap-%E3%81%8C%E5%88%86%E3%81%8B%E3%82%89%E3%81%AA%E3%81%84