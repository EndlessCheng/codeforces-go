package main

import (
	. "fmt"
	"io"
	"math/big"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var w, h, n, l, r int64
	Fscan(in, &w, &h, &n)
	l, r = 1, 1e18
	for l < r {
		m := (l + r) >> 1
		// or float64
		if new(big.Int).Mul(big.NewInt(m/w), big.NewInt(m/h)).Cmp(big.NewInt(n)) >= 0 {
			r = m
		} else {
			l = m + 1
		}
	}
	Fprint(out, l)
}

func main() { run(os.Stdin, os.Stdout) }
