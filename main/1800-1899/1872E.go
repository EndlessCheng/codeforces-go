package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1872E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, q, op, l, r int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		pre := make([]int, n+1)
		for i := range a {
			Fscan(in, &a[i])
			pre[i+1] = pre[i] ^ a[i]
		}
		Fscan(in, &s)
		xor1 := 0
		for i, b := range s {
			if b == '1' {
				xor1 ^= a[i]
			}
		}
		for Fscan(in, &q); q > 0; q-- {
			Fscan(in, &op, &l)
			if op == 1 {
				Fscan(in, &r)
				xor1 ^= pre[r] ^ pre[l-1]
			} else if l == 0 {
				Fprint(out, pre[n]^xor1, " ")
			} else {
				Fprint(out, xor1, " ")
			}
		}
		Fprintln(out)
	}
}

//func main() { CF1872E(os.Stdin, os.Stdout) }
