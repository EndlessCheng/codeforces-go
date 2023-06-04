package main

import . "fmt"

// https://space.bilibili.com/206214
func CF1839E() {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var n, s, fi, si int
	Scan(&n)
	a := make([]int, n+1)
	f := make([]int, n*300+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		Scan(&a[i])
		s += a[i]
		for j := s; j >= a[i]; j-- {
			if f[j] == 0 && f[j-a[i]] > 0 {
				f[j] = i
			}
		}
	}
	
	if s%2 > 0 || f[s/2] == 0 {
		Println("First")
		for {
			for i, v := range a {
				if v > 0 {
					fi = i
					break
				}
			}
			Println(fi)
			Scan(&si)
			if si <= 0 {
				break
			}
			v := min(a[fi], a[si])
			a[fi] -= v
			a[si] -= v
		}
	} else {
		left := make([]bool, n+1)
		for s /= 2; s > 0; s -= a[f[s]] {
			left[f[s]] = true
		}
		Println("Second")
		for {
			Scan(&fi)
			if fi <= 0 {
				break
			}
			for i, v := range a {
				if v > 0 && left[i] != left[fi] {
					si = i
					break
				}
			}
			v := min(a[fi], a[si])
			a[fi] -= v
			a[si] -= v
			Println(si)
		}
	}
}

//func main() { CF1839E() }
