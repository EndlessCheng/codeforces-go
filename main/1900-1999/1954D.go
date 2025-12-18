package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1954D(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, s, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		s += a[i]
	}
	slices.Sort(a)

	f := make([]int, s+1)
	f[0] = 1
	s = 0
	for _, v := range a {
		sumF := 0
		for _, fv := range f[:v] {
			sumF += fv
		}
		ans += sumF % mod * v

		for j := v; j <= s; j++ {
			ans += (j + v + 1) / 2 * f[j]
		}
		ans %= mod

		s += v
		for j := s; j >= v; j-- {
			f[j] = (f[j] + f[j-v]) % mod
		}
	}
	Fprint(out, ans)
}

//func main() { cf1954D(bufio.NewReader(os.Stdin), os.Stdout) }
