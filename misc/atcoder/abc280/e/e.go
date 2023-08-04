package main

import (
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	const mod = 998244353
	const inv100 = 828542813
	var n, p int
	Fscan(in, &n, &p)
	f := make([]int, n+1)
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = (p*f[i-2] + (100-p)*f[i-1] + 100) % mod * inv100 % mod
	}
	Fprint(out, f[n])
}

func main() { run(os.Stdin, os.Stdout) }
