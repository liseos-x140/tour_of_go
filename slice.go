package main

import (
	"math"

	"golang.org/x/tour/pic"
)

// Pic は任意の関数でuint8の画像生成の元となる配列を生成する関数
func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for i := range s {
		t := make([]uint8, dx)
		for j := range t {
			//t[j] = uint8((i + j) / 2)  // (x+y)/2
			//t[j] = uint8(i * j)  // x*y
			t[j] = uint8(math.Pow(float64(j), float64(i))) // x^y
		}
		s[i] = t
	}
	return s
}

func main() {
	pic.Show(Pic)
}
