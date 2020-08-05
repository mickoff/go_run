package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	d := float64(100)
	z := float64(1)
	i := 1
	for ; float64(0.001) < d; i++ {
		d = z
		z -= (z*z - x) / (2*z)
		d = math.Abs(d-z)
	}
	fmt.Println("誤差", d)
	fmt.Println("ループ回数", i)
	return z
}

func main() {
	fmt.Println(Sqrt(100))
}