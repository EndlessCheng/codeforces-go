package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF785D(in io.Reader, out io.Writer) {
	const mod int64 = 1e9 + 7
	pow := func(x, n int64) (res int64) {
		res = 1
		for ; n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	var s string
	Fscan(bufio.NewReader(in), &s)
	n := len(s)
	F := make([]int64, n+1)
	F[0] = 1
	for i := 1; i <= n; i++ {
		F[i] = F[i-1] * int64(i) % mod
	}
	invF := make([]int64, n+1)
	invF[n] = pow(F[n], mod-2)
	for i := n; i > 0; i-- {
		invF[i-1] = invF[i] * int64(i) % mod
	}
	C := func(n, k int) int64 { return F[n] * invF[k] % mod * invF[n-k] % mod }

	ans := int64(0)
	l, r := 0, strings.Count(s, ")")
	for _, b := range s {
		if r == 0 {
			break
		}
		if b == '(' { // ∑ 左边 i-1 个括号，右边 i 个括号
			l++
			ans = (ans + C(l+r-1, l)) % mod
		} else {
			r--
		}
	}
	Fprint(out, ans)
}

//func main() { CF785D(os.Stdin, os.Stdout) }
