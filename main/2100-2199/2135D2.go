package main

import . "fmt"

// https://github.com/EndlessCheng
func cf2135D2() {
	q := func(a []any) (r int) {
		Print("? ", len(a), " ")
		Println(a...)
		Scan(&r)
		return
	}

	const B = 123
	const C = 8257
	b := [C]any{}
	for i := range b {
		b[i] = B
	}
	a1 := [B * B]any{}
	for i := range a1 {
		a1[i] = 1
	}

	var T int
	for Scan(&T); T > 0; T-- {
		r := q(b[:])
		if r == 0 {
			r = q(a1[:])
			Println("!", (B*B-1)/r+1)
			continue
		}
		mn := ((C-1)/r + 1) * B
		mx := min(((C-1)/(r-1)+1)*B-1, 1e5)
		a := make([]any, 0, (mx-mn)*2)
		for i := 1; i <= mx-mn; i++ {
			a = append(a, mn, i)
		}
		Println("!", mx*2-mn-q(a))
	}
}

//func main() { cf2135D2() }
