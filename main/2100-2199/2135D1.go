package main

import (
	. "fmt"
)

// https://github.com/EndlessCheng
func cf2135D1() {
	q := func(a []any) (r int) {
		Print("? ", len(a), " ")
		Println(a...)
		Scan(&r)
		return
	}

	const n int = 1e5
	a1 := [n]any{}
	for i := range a1 {
		a1[i] = 1
	}

	var T int
	for Scan(&T); T > 0; T-- {
		r := q(a1[:])
		if r == 1 {
			Println("!", n)
			continue
		}
		mn, mx := (n-1)/r+1, (n-1)/(r-1)
		if mn == mx {
			Println("!", mn)
			continue
		}
		a := make([]any, 0, mn*2)
		for i := 1; i <= mn; i++ {
			a = append(a, mn, i)
		}
		Println("!", mn*3-q(a))
	}
}

//func main() { cf2135D1() }
