package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if (x < 0) {
		return 0, ErrNegativeSqrt(x)
	}
	d := float64(100)
	z := float64(1)
	i := 1
	for ; float64(0.001) < d; i++ {
		d = z
		z -= (z*z - x) / (2*z)
		d = math.Abs(d-z)
	}
	// fmt.Println("誤差", d)
	// fmt.Println("ループ回数", i)
	return z, nil
}

func main() {
	if i, err := Sqrt(2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
	if i, err := Sqrt(-2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
}