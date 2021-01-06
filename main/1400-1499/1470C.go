package main

import (
	. "fmt"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func CF1470C() {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var n, k, v int
	Scan(&n, &k)
	step := int(math.Sqrt(float64(n)))
	for gcd(step, n) > 1 {
		step++
	}
	for i := 0; ; i += step {
		Println("?", i%n+1)
		Scan(&v)
		if v == k {
			continue
		}
		if v < k {
			for j := i + 1; ; j++ {
				Println("?", j%n+1)
				Scan(&v)
				if v == k {
					Println("!", j%n+1)
					return
				}
			}
		} else {
			for j := i - 1; ; j-- {
				Println("?", (j%n+n)%n+1)
				Scan(&v)
				if v == k {
					Println("!", (j%n+n)%n+1)
					return
				}
			}
		}
	}
}

//func main() { CF1470C() }
