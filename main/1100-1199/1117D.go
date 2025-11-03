package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
const mod17 = 1_000_000_007

func kitamasa17(coef, a []int, n int) (ans int) {
	defer func() { ans = (ans%mod17 + mod17) % mod17 }()
	if n < len(a) {
		return a[n]
	}

	k := len(coef)
	compose := func(a, b []int) []int {
		c := make([]int, k)
		for _, v := range a {
			for j, w := range b {
				c[j] = (c[j] + v*w) % mod17
			}
			bk1 := b[k-1]
			for j := k - 1; j > 0; j-- {
				b[j] = (b[j-1] + bk1*coef[j]) % mod17
			}
			b[0] = bk1 * coef[0] % mod17
		}
		return c
	}

	resC := make([]int, k)
	resC[0] = 1
	c := make([]int, k)
	c[1] = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			resC = compose(c, resC)
		}
		c = compose(c, slices.Clone(c))
	}

	for i, c := range resC {
		ans = (ans + c*a[i]) % mod17
	}
	return
}

func cf1117D(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, m)
	for i := range a {
		a[i] = 1
	}
	coef := make([]int, m)
	coef[0] = 1
	coef[m-1] = 1
	Fprint(out, kitamasa17(coef, a, n))
}

//func main() { cf1117D(os.Stdin, os.Stdout) }
