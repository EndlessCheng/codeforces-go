package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, m, sz int
	Fscan(in, &n, &m, &sz)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	f := make([]int, m)
	for i := 1; i < m; i++ {
		f[i] = 1e9
	}
	for i := range sz {
		nf := make([]int, m)
		for j := range nf {
			nf[j] = 1e9
		}
		for tar := range m {
			op := 0
			for j := i; j < n; j += sz {
				op += (tar - a[j] + m) % m
			}
			for j := range m {
				nf[j] = min(nf[j], f[(j+tar)%m]+op)
			}
		}
		f = nf
	}
	Fprint(out, f[0])
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
