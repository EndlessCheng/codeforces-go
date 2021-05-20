package main

import . "fmt"

// github.com/EndlessCheng/codeforces-go
func CF896B() {
	var n, c, p, i int
	Scan(&n, &c, &c)
	a := make([]int, n+1)
	for l := n; l > 0; {
		if Scan(&p); p > c/2 {
			for i = n; a[i] >= p; i-- {
			}
		} else {
			for i = 1; 0 < a[i] && a[i] <= p; i++ {
			}
		}
		Println(i)
		if a[i] == 0 {
			l--
		}
		a[i] = p
	}
}

//func main() { CF896B() }
