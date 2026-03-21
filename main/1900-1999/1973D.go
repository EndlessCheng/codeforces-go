package main

import . "fmt"

// https://github.com/EndlessCheng
func q73(l, x int) int {
	Println("?", l, x)
	Scan(&x)
	return x
}

func cf1973D() {
	var T, n, k int
o:
	for Scan(&T); T > 0; T-- {
		Scan(&n, &k)

		m := -1
		for i := n; i > 0; i-- {
			if q73(1, i*n) == n {
				m = i
				break
			}
		}

		for j := n / k; j > 0; j-- {
			l := 1
			for i := range k {
				if l > n {
					break
				}
				l = q73(l, j*m) + 1
				if i == k-1 && l == n+1 {
					Println("!", j*m)
					Scan(&n)
					continue o
				}
			}
		}

		Println("! -1")
		Scan(&n)
	}
}

//func main() { cf1973D() }
