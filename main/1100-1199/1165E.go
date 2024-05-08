package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func cf1165E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 998244353
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		a[i] *= (n - i) * (i + 1)
	}
	slices.SortFunc(a, func(a, b int) int { return b - a })
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
	}
	slices.Sort(b)

	for i, v := range a {
		ans = (ans + v%mod*b[i]) % mod
	}
	Fprint(out, ans)
}

//func main() { cf1165E(os.Stdin, os.Stdout) }
