package main

import "fmt"

func fibonacci() func(int) int {
	f 	:= 0
	f_1 := 0
	f_2 := 0
	return func(x int) int {
		switch x {
		case 0:
			break
		case 1:
			f = 1
		default:
			f = f_1 + f_2
		}
		f_2 = f_1
		f_1 = f
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}