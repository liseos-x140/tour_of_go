package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount は文字列として与えられた引数から単語数をカウントする関数
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, word := range strings.Fields(s) {
		m[word]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
