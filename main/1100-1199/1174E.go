package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1174E(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n, ans int
	Fscan(in, &n)

	f := func(a []int) {
		res, fac := 1, 1
		for i, v := range a[:len(a)-1] {
			c := n/v - n/a[i+1]
			res = res * c % mod
			for range c - 1 {
				res = res * fac % mod
				fac++
			}
			fac++
		}
		ans += res
	}

	a := []int{}
	for v := 1; v <= n; v *= 2 {
		a = append(a, v)
	}
	f(a)

	m := len(a)
	if a[m-1]/2*3 <= n {
		for i := 1; i < m; i++ {
			b := slices.Clone(a)
			for j := i; j < m; j++ {
				b[j] = b[j] / 2 * 3
			}
			f(b)
		}
	}

	Fprint(out, ans%mod)
}

//func main() { cf1174E(os.Stdin, os.Stdout) }
