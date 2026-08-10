package main

import (
	"fmt"
	"math"
	"math/rand"
)

// 原题 https://codeforces.com/problemset/problem/1215/D

// github.com/EndlessCheng/codeforces-go
func sumGame(num string) bool {
	m := len(num) / 2
	d := 0
	for i, ch := range num {
		x := 9
		if ch != '?' {
			x = int(ch-'0') * 2
		}
		if i < m {
			d += x
		} else {
			d -= x
		}
	}
	return d != 0
}

func main() {
	bobWin := 0
	const n = 300
	const t = 1000000
	s := make([]byte, n)
	for range t {
		for i := range s {
			v := rand.Intn(11)
			if v == 10 {
				s[i] = '?'
			} else {
				s[i] = '0' + byte(v)
			}
		}
		if !sumGame(string(s)) {
			bobWin++
		}
	}
	fmt.Println(float64(bobWin) / t)
	fmt.Println(1/ math.Sqrt(float64(n * 60) * math.Pi))
}
