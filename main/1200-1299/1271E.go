package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1271E(in io.Reader, out io.Writer) {
	max := func(a, b uint64) uint64 {
		if a > b {
			return a
		}
		return b
	}

	var n, k uint64
	Fscan(in, &n, &k)
	f := func(k uint64) uint64 {
		b := uint64(1) << (bits.Len64(k) - 1)
		return (n - k + b) / b
	}
	Fprint(out, max(f(k), f(k+1)*2))
}

//func main() { CF1271E(os.Stdin, os.Stdout) }
