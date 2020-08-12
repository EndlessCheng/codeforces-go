package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func CF1033D(_r io.Reader, _w io.Writer) {
	const mod int64 = 998244353
	sqrt := func(a int64) int64 {
		r := int64(math.Round(math.Sqrt(float64(a))))
		if r*r == a {
			return r
		}
		return -1
	}
	cbrt := func(a int64) int64 {
		r := int64(math.Round(math.Cbrt(float64(a))))
		if r*r*r == a {
			return r
		}
		return -1
	}
	gcd := func(a, b int64) int64 {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int64, 0, n)
	vis := make([]bool, 0, n)
	cntv := map[int64]int{}
	cnte := map[int64]int{}
	for ; n > 0; n-- {
		var v int64
		Fscan(in, &v)
		ok := true
		if r := cbrt(v); r != -1 {
			cnte[r] += 3
		} else if r := sqrt(v); r != -1 {
			if r4 := sqrt(r); r4 != -1 {
				cnte[r4] += 4
			} else {
				cnte[r] += 2
			}
		} else {
			ok = false
		}
		cntv[v]++
		if cntv[v] == 1 {
			a = append(a, v)
			vis = append(vis, ok)
		}
	}

	ans := int64(1)
	for i, x := range a {
		if !vis[i] {
			for j, y := range a {
				if j == i {
					continue
				}
				if g := gcd(x, y); g > 1 {
					cnte[g] += cntv[x]
					cnte[x/g] += cntv[x]
					if !vis[j] {
						cnte[g] += cntv[y]
						cnte[y/g] += cntv[y]
						vis[j] = true
					}
					vis[i] = true
					break
				}
			}
		}
		if !vis[i] {
			c := cntv[x] + 1
			ans = ans * int64(c*c) % mod
		}
	}
	for _, e := range cnte {
		ans = ans * int64(e+1) % mod
	}
	Fprintln(_w, ans)
}

//func main() { CF1033D(os.Stdin, os.Stdout) }
