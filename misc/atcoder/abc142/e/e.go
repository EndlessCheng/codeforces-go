package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, m, k, v int
	Fscan(in, &n, &m)
	a := make([]struct{ a, msk int }, m)
	for i := range a {
		Fscan(in, &a[i].a)
		for Fscan(in, &k); k > 0; k-- {
			Fscan(in, &v)
			a[i].msk |= 1 << (v - 1)
		}
	}
	f := make([]int, 1<<n)
	for i := range f {
		f[i] = 1e9
	}
	f[0] = 0
	for i, fi := range f {
		for _, p := range a {
			f[i|p.msk] = min(f[i|p.msk], fi+p.a)
		}
	}
	if f[1<<n-1] == 1e9 {
		Fprint(out, -1)
	} else {
		Fprint(out, f[1<<n-1])
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
