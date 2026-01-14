package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1437F(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)

	f := make([]int, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		p := 1
		j := i - 2
		for j >= 0 && a[j]*2 > a[i-1] {
			p = p * (n - j - 2) % mod
			j--
		}
		f[i] = (f[i-1]*(n-i) + f[j+1]*p) % mod
	}
	Fprint(out, f[n])
}

//func main() { cf1437F(bufio.NewReader(os.Stdin), os.Stdout) }
