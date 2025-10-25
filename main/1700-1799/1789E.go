package main

import (
	. "fmt"
	"io"
)

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
				// s = Ax+B(x+1)
				// s = Qx+R，其中 Q>=R
				// Q = s/x, R = s%x
				// s/x >= s%x
				// 枚举 Q，那么满足 s/x=Q 的 s 的范围为 [Qx,Qx+x-1]
				// 此外 s%x<=Q，s 的范围为 [Qx,Qx+Q]
				// 二者结合，所以 s 的范围为 [Qx, Qx+min(x-1,Q)] （当然，上界不能超过 mx）
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
