package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1081C(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, m, k int
	Fscan(in, &n, &m, &k)
	f := make([]int64, k+1)
	f[0] = int64(m)
	for i := 1; i < n; i++ {
		for j := k; j > 0; j-- {
			f[j] = (f[j] + f[j-1]*int64(m-1)) % mod
		}
	}
	Fprint(out, f[k])
}

//func main() { CF1081C(os.Stdin, os.Stdout) }
