package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1542E1(in io.Reader, out io.Writer) {
	var n, mod, ans int
	Fscan(in, &n, &mod)

	maxK := n * (n - 1) / 2
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, maxK+1)
	}
	f[0][0] = 1
	sum := make([]int, maxK+2)
	for i := 1; i <= n; i++ {
		mx := i*(i-1)/2 + 1
		for j, v := range f[i-1][:mx] {
			sum[j+1] = sum[j] + v
		}
		for j := range mx {
			f[i][j] = (sum[j+1] - sum[max(j-i+1, 0)]) % mod
		}
	}

	perm := 1
	for i := n; i > 3; i-- { // 剩余 i 个数：此处枚举 P 和 Q 选什么
		row := f[i-1]
		res := 0
		for delta := 1; delta < i; delta++ { // delta = Q 新增逆序对 - P 新增逆序对，有 i-delta 个方案
			s := 0
			for j := delta + 1; j <= maxK; j++ { // 枚举 row 中的 P 的逆序对 j，累加 Q 的逆序对 s
				s = (s + row[j-delta-1]) % mod
				res = (res + row[j]*s%mod*(i-delta)) % mod
			}
		}
		ans = (ans + perm*res) % mod
		perm = perm * i % mod
	}
	Fprint(out, ans)
}

//func main() { cf1542E1(os.Stdin, os.Stdout) }
