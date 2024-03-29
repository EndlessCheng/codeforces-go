package main

import . "fmt"

// github.com/EndlessCheng/codeforces-go
func CF1355F() {
	const mx = 1e3
	ps := []int{}
	v := [mx]bool{}
	for i := 2; i < mx; i++ {
		if !v[i] {
			ps = append(ps, i)
			for j := 2 * i; j < mx; j += i {
				v[j] = true
			}
		}
	}
	const upper int64 = 1e18

	var t int
	var g int64
	for Scan(&t); t > 0; t-- {
		ans := 1
		es := [mx]int{}
		for i := 0; i < 22; i++ {
			checks := []int64{}
			q := int64(1)
		o:
			for _, p := range ps {
				if e := es[p]; e >= 0 {
					for p := int64(p); e >= 0; e-- {
						if q > upper/p {
							break o
						}
						q *= p
					}
					checks = append(checks, int64(p))
				}
			}
			Println("?", q)
			Scan(&g)
			for _, p := range checks {
				e := 0
				for ; g%p == 0; g /= p {
					e++
				}
				if e > es[p] {
					es[p] = e
					if i == 21 {
						ans *= e + 1
					}
				} else {
					es[p] = -1
					ans *= e + 1
				}
			}
		}
		ans *= 2
		if ans < 8 {
			ans = 8
		}
		Println("!", ans)
	}
}

//func main() { CF1355F() }
