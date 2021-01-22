package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353
	const mod2 = (mod + 1) / 2

	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	m := append([]int(nil), a...)
	s := make([]int, k+1)
	s[0] = n
	for x := 1; x <= k; x++ {
		for _, v := range m {
			s[x] += v
		}
		s[x] %= mod
		for i, v := range a {
			m[i] = m[i] * v % mod
		}
	}

	c := make([][]int, k+1)
	c[0] = []int{1}
	for x, p2 := 1, 2; x <= k; x++ {
		c[x] = make([]int, x+1)
		c[x][0], c[x][x] = 1, 1
		for j := 1; j < x; j++ {
			c[x][j] = (c[x-1][j-1] + c[x-1][j]) % mod
		}
		sum := -p2 * s[x] % mod
		for i, v := range c[x] {
			sum += v * s[i] % mod * s[x-i] % mod
		}
		Fprintln(out, (sum%mod+mod)*mod2%mod)
		p2 = p2 * 2 % mod
	}
}

func main() { run(os.Stdin, os.Stdout) }
