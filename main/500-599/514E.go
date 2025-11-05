package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
const mod14 = 1_000_000_007

func pow14(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod14
		}
		x = x * x % mod14
	}
	return res
}

func berlekampMassey14(a []int) (coef []int) {
	var preC []int
	preI, preD := -1, 0
	for i, v := range a {
		d := v
		for j, c := range coef {
			d = (d - c*a[i-1-j]) % mod14
		}
		if d == 0 {
			continue
		}

		if preI < 0 {
			coef = make([]int, i+1)
			preI, preD = i, d
			continue
		}

		bias := i - preI
		oldLen := len(coef)
		newLen := bias + len(preC)
		var tmp []int
		if newLen > oldLen {
			tmp = slices.Clone(coef)
			coef = slices.Grow(coef, newLen-oldLen)[:newLen]
		}

		delta := d * pow14(preD, mod14-2) % mod14
		coef[bias-1] = (coef[bias-1] + delta) % mod14
		for j, c := range preC {
			coef[bias+j] = (coef[bias+j] - delta*c) % mod14
		}

		if newLen > oldLen {
			preC = tmp
			preI, preD = i, d
		}
	}

	return
}

func kitamasa14(coef, a []int, n int) (ans int) {
	defer func() { ans = (ans%mod14 + mod14) % mod14 }()
	if n < len(a) {
		return a[n]
	}

	k := len(coef)
	if k == 0 {
		return
	}
	if k == 1 {
		return a[0] * pow14(coef[0], n)
	}

	compose := func(a, b []int) []int {
		c := make([]int, k)
		for _, v := range a {
			for j, w := range b {
				c[j] = (c[j] + v*w) % mod14
			}
			bk1 := b[k-1]
			for j := k - 1; j > 0; j-- {
				b[j] = (b[j-1] + bk1*coef[j]) % mod14
			}
			b[0] = bk1 * coef[0] % mod14
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
		ans = (ans + c*a[i]) % mod14
	}
	return
}

func cf514E(in io.Reader, out io.Writer) {
	var n, x, v int
	Fscan(in, &n, &x)
	const mx = 100
	cnt := [mx + 1]int{}
	for range n {
		Fscan(in, &v)
		cnt[v]++
	}

	a := make([]int, mx*2+2)
	a[0] = 1
	for i := 1; i < len(a); i++ {
		a[i] = 1
		for j, c := range cnt[1 : min(i, mx)+1] {
			a[i] += c * a[i-1-j]
		}
		a[i] %= mod14
	}
	coef := berlekampMassey14(a)
	slices.Reverse(coef)
	Fprint(out, kitamasa14(coef, a, x))
}

//func main() { cf514E(bufio.NewReader(os.Stdin), os.Stdout) }
