package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1444B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int64 = 998244353
	pow := func(x, n int64) (res int64) {
		res = 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	var n int64
	Fscan(in, &n)
	a := make([]int, 2*n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	s := int64(0)
	for _, v := range a[:n] {
		s -= int64(v)
	}
	for _, v := range a[n:] {
		s += int64(v)
	}
	f := int64(1)
	for i := int64(2); i <= n; i++ {
		f = f * i % mod
	}
	inv := pow(f, mod-2)
	for i := n + 1; i <= 2*n; i++ {
		f = f * i % mod
	}
	Fprint(out, s%mod*f%mod*inv%mod*inv%mod)
}

//func main() { CF1444B(os.Stdin, os.Stdout) }
