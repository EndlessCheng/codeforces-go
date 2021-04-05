package main

import . "fmt"

// github.com/EndlessCheng/codeforces-go
func CF1503B() {
	var n, a int
	Scan(&n)
	for i, p, q, m := 0, 0, 1, n*n; i < m; i++ {
		Scan(&a)
		if q >= m || a > 1 && p < m {
			c := p%n + 1
			if n&1 == 0 && p/n&1 > 0 {
				c++
			}
			if a == 1 {
				Println(3, p/n+1, c)
			} else {
				Println(1, p/n+1, c)
			}
			p += 2
		} else {
			c := q%n + 1
			if n&1 == 0 && q/n&1 > 0 {
				c--
			}
			if a == 2 {
				Println(3, q/n+1, c)
			} else {
				Println(2, q/n+1, c)
			}
			q += 2
		}
	}
}

//func main() { CF1503B() }
