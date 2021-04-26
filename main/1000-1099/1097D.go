package main

import (
	. "fmt"
	"io"
)

// 参考 https://www.luogu.com.cn/blog/ouuan/solution-cf1097d

// github.com/EndlessCheng/codeforces-go
func CF1097D(in io.Reader, out io.Writer) {
	const mod int64 = 1e9 + 7
	inv := [51]int64{1: 1}
	for i := int64(2); i < 51; i++ {
		inv[i] = (mod - mod/i) * inv[mod%i] % mod
	}
	var n int64
	var k int
	Fscan(in, &n, &k)
	ans := int64(1)
	do := func(p int64, e int) {
		f := [51]int64{}
		f[e] = 1
		for i := 0; i < k; i++ {
			for j := e; j >= 0; j-- {
				f[j] = (f[j+1] + f[j]*inv[j+1]) % mod
			}
		}
		sum, pp := int64(0), int64(1)
		for _, v := range f[:] {
			sum = (sum + pp*v) % mod
			pp = pp * p % mod
		}
		ans = ans * sum % mod
	}
	for i := int64(2); i*i <= n; i++ {
		e := 0
		for ; n%i == 0; n /= i {
			e++
		}
		if e > 0 {
			do(i, e)
		}
	}
	if n > 1 {
		do(n, 1)
	}
	Fprint(out, ans)
}

//func main() { CF1097D(os.Stdin, os.Stdout) }
