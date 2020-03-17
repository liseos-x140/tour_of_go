package main

import (
	"fmt"
)

// Sqrt はニュートン法を用いた平方根の計算を行う関数です．
func Sqrt(x float64) float64 {
	z := x
	i := 0

	for {
		// ニュートン法計算
		d := (z*z - x) / (2 * z)
		z -= d
		i++ // counter

		// 変化量が少ない場合 return
		if d*d < 1e-15 {
			fmt.Printf("loop num: %d\n", i)
			return z
		}
	}
}

func main() {
	fmt.Println(Sqrt(119))
}
