package main

import . "fmt"

// https://github.com/EndlessCheng
func cf1493F() {
	var n, m, x int
	Scan(&n, &m)
	n0, m0 := n, m
	ans := 1

	q := func(tp, size, mx int) int {
		for i := size; i < mx; i *= 2 {
			if tp == 1 {
				Println("?", n0, min(i, mx-i), 1, 1, 1, i+1)
			}
			if tp == 2 {
				Println("?", min(i, mx-i), m0, 1, 1, i+1, 1)
			}
			Scan(&x)
			if x == 0 {
				return 0
			}
		}
		return 1
	}

	f := func(n, n0, tp int) {
		for i := 2; i <= n; i++ {
			val, now := 1, 1
			for n%i == 0 {
				n /= i
				val *= i
			}
			for j := val / i; j >= 1; j /= i {
				if q(tp, n0/val*j, n0/val*j*i) != 0 {
					now++
				} else {
					break
				}
			}
			ans *= now
		}
	}
	f(n, n, 2)
	f(m, m, 1)
	Println("!", ans)
}

//func main() { cf1493F() }
