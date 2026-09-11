package main

import . "fmt"

// https://github.com/EndlessCheng
func cf1291F() {
	var n, k int
	var s string
	Scan(&n, &k)
	mod := n / k
	f := make([]int, n)
	for i := range f {
		f[i] = 1
	}
	for x := range mod {
		i, d := x, 1
		for range mod {
			for j := range k {
				if f[i*k+j] != 0 {
					Println("?", i*k+j+1)
					Scan(&s)
					if s[0] == 'Y' {
						f[i*k+j] = 0
					}
				}
			}
			i = ((i+d)%mod + mod) % mod
			if d > 0 {
				d = -d - 1
			} else {
				d = -d + 1
			}
		}
		Println("R")
	}

	ans := 0
	for _, v := range f {
		ans += v
	}
	Println("!", ans)
}

//func main() { cf1291F() }
