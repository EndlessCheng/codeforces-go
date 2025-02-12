package main

import (
	. "fmt"
	"io"
	"os"
	"slices"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, k int
	Fscan(in, &n, &k)
	f := []int{1}
	for i := 1; i <= n; i++ {
		f = slices.Insert(f, 0, 0)
		for j := i / 2; j > 0; j-- {
			f[j] = (f[j] + f[j*2]) % mod
		}
	}
	Fprint(out, f[k])
}

func main() { run(os.Stdin, os.Stdout) }
