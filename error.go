package main

import (
	"fmt"
)

// ErrNegativeSqrt Sqrt 関数に負の値が与えられたときのカスラムエラー型
type ErrNegativeSqrt float64

// ErrNegativeSqrt のエラー出力
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt はニュートン法を用いた平方根の計算を行う関数 + Error
func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		z := x
		i := 0

		for {
			// ニュートン法計算
			d := (z*z - x) / (2 * z)
			z -= d
			i++ // counter

			// 変化量が少ない場合 return
			if d*d < 1e-15 {
				// fmt.Printf("loop num: %d\n", i)
				return z, nil
			}
		}
	}
	return 0, ErrNegativeSqrt(x)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
