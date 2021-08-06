package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF757D(in io.Reader, out io.Writer) {
	const mod int = 1e9 + 7
	const mx = 20
	var n, ans int
	var s []byte
	Fscan(in, &n, &s)
	for i := range s {
		s[i] &= 1
	}
	// 注意，不要写 make([][1<<mx]int, n)，由于 Go 内存管理的原因，会将申请的内存塞到一个不小于所需空间的 2 的幂次的内存块中
	// 在 n=75 时，所需空间为 300MB，因此会塞到 512MB 大小的内存中，但是申请这部分内存又会导致 MLE
	// 下面的写法，每次只申请 4MB 的内存，最后实际使用内存约为 334MB https://codeforces.com/contest/757/submission/124971542
	// 另一种写法是声明一个全局变量 var dp [75][1<<mx]int，最后实际使用内存约为 311MB https://codeforces.com/contest/757/submission/124969909
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 1<<mx)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(p, mask int) (res int) {
		if mask&(mask+1) == 0 && mask > 0 {
			res = 1
		}
		if p == n {
			return
		}
		dv := &dp[p][mask]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		for j, v := p, byte(0); j < n; j++ {
			v = v<<1 | s[j]
			if v > mx {
				break
			}
			if v > 0 {
				res = (res + f(j+1, mask|1<<(v-1))) % mod
			}
		}
		return
	}
	for i := range s {
		ans = (ans + f(i, 0)) % mod
	}
	Fprint(out, ans)
}

//func main() { CF757D(os.Stdin, os.Stdout) }
