package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1391C(in io.Reader, out io.Writer) {
	const mod int64 = 1e9 + 7
	var n, f, p int64 = 0, 1, 1
	Fscan(in, &n)
	for ; n > 1; n-- {
		f = f * n % mod
		p = p * 2 % mod
	}
	Fprint(out, (f-p+mod)%mod)
}

//func main() { CF1391C(os.Stdin, os.Stdout) }
