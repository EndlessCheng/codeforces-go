package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1324E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n, h, l, r int
	Fscan(in, &n, &h, &l, &r)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	f := [2][]int{make([]int, h), make([]int, h)}
	for s := l; s <= r; s++ {
		f[n&1][s] = 1
	}
	for i := n - 1; i >= 0; i-- {
		for s := 0; s < h; s++ {
			f[i&1][s] = max(f[i&1^1][(s+a[i])%h], f[i&1^1][(s+a[i]-1)%h])
			if i > 0 && l <= s && s <= r {
				f[i&1][s]++
			}
		}
	}
	Fprint(out, f[0][0])
}

//func main() { CF1324E(os.Stdin, os.Stdout) }
