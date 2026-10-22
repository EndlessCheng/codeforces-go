package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1349F1(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353
	var n int
	Fscan(in, &n)

	f := make([]int, n+1)
	g := make([]int, n+1)
	f[1] = 1
	g[1] = 1
	for i := 2; i <= n; i++ {
		nf := make([]int, n+1)
		ng := make([]int, n+1)
		for k := 1; k <= i; k++ {
			nf[k] = (k*f[k] + (i-k+1)*f[k-1]) % mod
			ng[k] = (i*g[k] + nf[k]) % mod
		}
		f = nf
		g = ng
	}

	for k := 1; k <= n; k++ {
		Fprint(out, g[k], " ")
	}
	Fprintln(out)
}

//func main() { cf1349F1(bufio.NewReader(os.Stdin), os.Stdout) }
