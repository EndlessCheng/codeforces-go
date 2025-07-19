package main

import (
	. "fmt"
	"io"
)

// 题解 https://www.luogu.com.cn/article/gamc2rrj

// https://github.com/EndlessCheng
func cf1526E(in io.Reader, out io.Writer) {
	const mod = 998244353
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	comb := func(n, k int) int {
		if n < k {
			return 0
		}
		k = min(k, n-k)
		a, b := 1, 1
		for i := 1; i <= k; i++ {
			a = a * (n - i + 1) % mod
			b = b * i % mod
		}
		return a * pow(b, mod-2) % mod
	}

	var n, k int
	Fscan(in, &n, &k)
	sa := make([]int, n)
	rank := make([]int, n+1)
	rank[n] = -1 // 空串的排名
	for i := range sa {
		Fscan(in, &sa[i])
		rank[sa[i]] = i
	}
	for i := range n - 1 {
		if rank[sa[i]+1] > rank[sa[i+1]+1] {
			k--
		}
	}
	Fprint(out, comb(n+k-1, n))
}

//func main() { cf1526E(bufio.NewReader(os.Stdin), os.Stdout) }
