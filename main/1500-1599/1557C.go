package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1557C(in io.Reader, out io.Writer) {
	const mod int64 = 1e9 + 7
	pow := func(x int64, n int) (res int64) {
		res = 1
		for ; n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return
	}
	p2 := [2e5]int64{1}
	for i := 1; i < 2e5; i++ {
		p2[i] = p2[i-1] << 1 % mod
	}

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		if n&1 > 0 {
			Fprintln(out, pow(p2[n-1]+1, k))
		} else {
			ans := int64(1)
			for i := 0; i < k; i++ {
				ans = (pow(p2[i], n) + (p2[n-1]-1)*ans) % mod
			}
			Fprintln(out, ans)
		}
	}
}

//func main() { CF1557C(os.Stdin, os.Stdout) }
