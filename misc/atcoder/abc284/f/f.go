package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func calcZ(s string) []int {
	n := len(s)
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r {
			z[i] = min(z[i-l], r-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
	}
	z[0] = n
	return z
}

func run(in io.Reader, out io.Writer) {
	var n int
	var t string
	Fscan(bufio.NewReader(in), &n, &t)
	for i := 0; i < n; i++ {
		if t[i] != t[n*2-1-i] {
			goto next
		}
	}
	Fprintln(out, t[n:])
	Fprint(out, 0)
	return
next:
	rev := make([]byte, len(t))
	for i := range t {
		rev[len(t)-1-i] = t[i]
	}
	r := string(rev)
	z1 := calcZ(t + r)
	z2 := calcZ(r + t)
	for i := 1; i < n; i++ {
		if z1[n*3-i] >= i && z2[n*2+i] >= n-i {
			Fprintln(out, t[:i]+t[n+i:])
			Fprint(out, i)
			return
		}
	}
	Fprint(out, -1)
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int { if b < a { return b }; return a }
