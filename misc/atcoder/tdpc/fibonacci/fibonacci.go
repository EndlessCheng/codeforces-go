package main

import (
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func kitamasa(k, n int) (ans int) {
	const mod = 1_000_000_007
	if n < k {
		return 1
	}

	mul := func(a, b []int) []int {
		c := make([]int, k)
		for _, v := range a {
			for j, w := range b {
				c[j] = (c[j] + v*w) % mod
			}
			bk := b[k-1]
			for i := k - 1; i > 0; i-- {
				b[i] = (b[i-1] + bk) % mod
			}
			b[0] = bk
		}
		return c
	}

	resC := make([]int, k)
	resC[0] = 1
	c := make([]int, k)
	c[1] = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			resC = mul(c, resC)
		}
		c = mul(c, append([]int{}, c...))
	}

	for _, c := range resC {
		ans += c
	}
	return ans % mod
}

func run(in io.Reader, out io.Writer) {
	var k, n int
	Fscan(in, &k, &n)
	Fprintln(out, kitamasa(k, n-1))
}

func main() { run(os.Stdin, os.Stdout) }
