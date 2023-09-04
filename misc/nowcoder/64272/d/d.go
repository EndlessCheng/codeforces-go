package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 1_000_000_007
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for i, v := range a {
		ans = (ans + v*(v+1)/2) % mod
		l, r := i-1, i+1
		for l >= 0 && r < n && a[l] == a[r] {
			ans += a[l]
			l--
			r++
		}
		if l >= 0 && r < n {
			ans += min(a[l], a[r])
		}
	}
	Fprint(out, ans%mod)
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
