package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2071D2(in io.Reader, out io.Writer) {
	var T, n, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &l, &r)
		a := make([]int, n*2+3)
		pre := make([]int, n*2+3)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			pre[i] = pre[i-1] + a[i]
		}
		if n%2 == 0 {
			n++
			a[n] = pre[n/2] % 2
			pre[n] = pre[n-1] + a[n]
		}
		for i := n + 1; i <= n*2; i++ {
			a[i] = pre[i/2] % 2
			pre[i] = pre[i-1] + a[i]
		}

		even := make([]int, n*2+1)
		for i := 1; i <= n*2; i++ {
			even[i] = even[i-1]
			if i%2 == 0 {
				even[i] += a[i]
			}
		}

		f := func(m int) (res int) {
			for {
				if m <= n*2 {
					res ^= a[m]
					break
				}
				res ^= pre[n] % 2
				m /= 2
				if m%2 > 0 {
					break
				}
			}
			return
		}
		var sum func(int) (int, int)
		sum = func(m int) (int, int) {
			if m <= n*2 {
				return even[m], pre[m] - even[m]
			}
			eve := even[n*2-1]
			odd := pre[n*2-1] - eve
			if m%2 == 0 {
				m++
				odd -= f(m)
			}
			if m/2%2 > 0 {
				m++
				eve -= f(m)
				m++
				odd -= f(m)
			}
			e, _ := sum(m / 2)
			e -= even[n]
			both := e
			if pre[n]%2 > 0 {
				both = m/2 - n + 1 - e
			}
			return eve + both, odd + both
		}

		re, ro := sum(r)
		le, lo := sum(l - 1)
		Fprintln(out, re+ro-le-lo)
	}
}

//func main() { cf2071D2(bufio.NewReader(os.Stdin), os.Stdout) }
