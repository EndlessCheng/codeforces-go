package main

import (
	. "fmt"
	"io"
)

// https://gemini.google.com/app/95e8b9d866c5ccd5

// https://github.com/EndlessCheng
func cf1789E(in io.Reader, out io.Writer) {
	const mod = 998244353
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		mx := a[n-1]
		s := make([]int32, mx+1)
		for _, v := range a {
			s[v] = 1
		}
		for i := range mx {
			s[i+1] += s[i]
		}

		ans := mx * n % mod
		for l, r := 1, 0; l < mx; l = r + 1 {
			cnt := 0
			x := mx / l
			if x == (mx-1)/l+1 {
				// 统计 x 的倍数
				for i := x; i <= mx; i += x {
					cnt += int(s[i] - s[i-1])
				}
			} else {
				q := mx / x
				for i := 1; i <= q; i++ {
					cnt += int(s[min(i*x+min(x-1, i), mx)] - s[i*x-1])
				}
			}
			r = min(mx/(mx/l), (mx-1)/((mx-1)/l))
			ans = (ans + (l+r)*(r-l+1)/2%mod*cnt) % mod
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1789E(bufio.NewReader(os.Stdin), os.Stdout) }
