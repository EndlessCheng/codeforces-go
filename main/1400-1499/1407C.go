package main

import . "fmt"

// https://github.com/EndlessCheng
func cf1407C() {
	q := func(i, j int) (m int) {
		Println("?", i, j)
		Scan(&m)
		return
	}
	var n int
	Scan(&n)
	p := make([]any, n+1)
	p[0] = "!"
	i := 1
	for j := 2; j <= n; j++ {
		m1 := q(i, j)
		m2 := q(j, i)
		if m1 > m2 {
			p[i] = m1
			i = j
		} else {
			p[j] = m2
		}
	}
	p[i] = n
	Println(p...)
}

//func main() { cf1407C() }
