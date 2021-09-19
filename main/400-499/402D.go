package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF402D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var n, m, v, g, ans int
	Fscan(in, &n, &m)
	a := make([]int, n)
	gs := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		g = gcd(g, a[i])
		gs[i] = g
	}
	bad := make(map[int]bool, m)
	for ; m > 0; m-- {
		Fscan(in, &v)
		bad[v] = true
	}

	f := func(x int) (res int) {
		k := bits.TrailingZeros(uint(x))
		if bad[2] {
			res -= k
		} else {
			res += k
		}
		x >>= k
		for i := 3; i*i <= x; i += 2 {
			for x%i == 0 {
				x /= i
				if bad[i] {
					res--
				} else {
					res++
				}
			}
		}
		if x > 1 {
			if bad[x] {
				res--
			} else {
				res++
			}
		}
		return
	}
	for _, v := range a {
		ans += f(v)
	}
	g = 1
	for i := n - 1; i >= 0; i-- {
		if res := f(gs[i] / g); res < 0 {
			ans -= (i + 1) * res
			g = gs[i]
		}
	}
	Fprint(out, ans)
}

//func main() { CF402D(os.Stdin, os.Stdout) }
