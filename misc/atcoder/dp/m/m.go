package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n, k, c int
	Fscan(in, &n, &k)
	f := make([]int, k+1)
	f[0] = 1
	for ; n > 0; n-- {
		Fscan(in, &c)
		for j := 1; j <= k; j++ {
			f[j] = (f[j] + f[j-1]) % mod
		}
		for j := k; j > c; j-- {
			f[j] -= f[j-c-1]
		}
	}
	Fprint(out, (f[k]%mod+mod)%mod)
}

func main() { run(os.Stdin, os.Stdout) }
