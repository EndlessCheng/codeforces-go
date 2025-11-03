package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
const mod95 = 1_000_000_007

func pow95(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod95
		}
		x = x * x % mod95
	}
	return res
}

func berlekampMassey95(a []int) (coef []int) {
	var preC []int
	preI, preD := -1, 0
	for i, v := range a {
		d := v
		for j, c := range coef {
			d = (d - c*a[i-1-j]) % mod95
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

		delta := d * pow95(preD, mod95-2) % mod95
		coef[bias-1] = (coef[bias-1] + delta) % mod95
		for j, c := range preC {
			coef[bias+j] = (coef[bias+j] - delta*c) % mod95
		}

		if newLen > oldLen {
			preC = tmp
			preI, preD = i, d
		}
	}

	return
}

func kitamasa95(coef, a []int, n int) (ans int) {
	defer func() { ans = (ans%mod95 + mod95) % mod95 }()
	if n < len(a) {
		return a[n]
	}

	k := len(coef)
	if k == 0 {
		return
	}
	if k == 1 {
		return a[0] * pow95(coef[0], n)
	}

	compose := func(a, b []int) []int {
		c := make([]int, k)
		for _, v := range a {
			for j, w := range b {
				c[j] = (c[j] + v*w) % mod95
			}
			bk1 := b[k-1]
			for j := k - 1; j > 0; j-- {
				b[j] = (b[j-1] + bk1*coef[j]) % mod95
			}
			b[0] = bk1 * coef[0] % mod95
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
		ans = (ans + c*a[i]) % mod95
	}
	return
}

func cf1895F(in io.Reader, out io.Writer) {
	var T, n, x, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x, &k)
		if n == 1 {
			Fprintln(out, k)
			continue
		}

		ans := (x + k) * pow95(k*2+1, n-1)
		if x == 0 {
			Fprintln(out, ans%mod95)
			continue
		}

		f := make([]int, x)
		for i := range f {
			f[i] = 1
		}
		sum := make([]int, x+1)

		a := make([]int, x*2)
		for i := range a {
			for j, v := range f {
				sum[j+1] = sum[j] + v
			}
			for j := range f {
				f[j] = (sum[min(j+k+1, x)] - sum[max(j-k, 0)]) % mod95
				a[i] += f[j]
			}
			a[i] %= mod95
		}

		coef := berlekampMassey95(a)
		slices.Reverse(coef)
		ans -= kitamasa95(coef, a, n-2)

		Fprintln(out, (ans%mod95+mod95)%mod95)
	}
}

//func main() { cf1895F(bufio.NewReader(os.Stdin), os.Stdout) }
