package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF414B(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var u, n, ans int
	Fscan(in, &u, &n)
	f := make([]int, u+1)
	f[1] = 1
	for i := 0; i < n; i++ {
		for j := u; j > 0; j-- {
			for k := j * 2; k <= u; k += j {
				f[k] = (f[k] + f[j]) % mod
			}
		}
	}
	for _, v := range f {
		ans = (ans + v) % mod
	}
	Fprint(out, ans)
}

//func main() { CF414B(os.Stdin, os.Stdout) }
