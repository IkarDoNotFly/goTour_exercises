package main

import (
	"fmt"
	"math"
)

//Exercise: Loops and Functions

func Sqrt(x float64) float64 {
	z:=45.0
	for ; z*z-x>0.0000001;{
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
	}
	return z
}


func main() {
	Sqrt(4)
	fmt.Println(math.Sqrt(4))
}
