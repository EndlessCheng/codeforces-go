package main

import . "fmt"

// https://github.com/EndlessCheng
func cf1290D() {
	var n, k int
	var s string
	Scan(&n, &k)
	ok := make([]bool, n)
	for i := range ok {
		ok[i] = true
	}

	m := n / k
	for x := range m {
		i, d := x, 1
		for range m {
			for j := range k {
				if ok[i*k+j] {
					Println("?", i*k+j+1)
					Scan(&s)
					if s[0] == 'Y' {
						ok[i*k+j] = false
					}
				}
			}
			i = ((i+d)%m + m) % m
			if d > 0 {
				d = -d - 1
			} else {
				d = -d + 1
			}
		}
		Println("R")
	}

	ans := 0
	for _, v := range ok {
		if v {
			ans++
		}
	}
	Println("!", ans)
}

//func main() { cf1290D() }
