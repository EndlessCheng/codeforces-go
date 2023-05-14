package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1787E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &x)
		xor, g := 0, 0
		for i := 1; i <= n; i++ {
			xor ^= i
			if i == x || i <= i^x && i^x <= n {
				g++
			}
		}
		if g < k || xor != k%2*x {
			Fprintln(out, "NO")
			continue
		}
		Fprintln(out, "YES")
		vis := make([]bool, n+1)
		left := n
		for i := 1; i <= n && k > 1; i++ {
			if i == x {
				Fprintln(out, 1, i)
				vis[i] = true
				left--
				k--
			} else if i^x <= n && !vis[i] {
				Fprintln(out, 2, i, i^x)
				vis[i] = true
				vis[i^x] = true
				left -= 2
				k--
			}
		}
		Fprint(out, left)
		for i := 1; i <= n; i++ {
			if !vis[i] {
				Fprint(out, " ", i)
			}
		}
		Fprintln(out)
	}
}

//func main() { CF1787E(os.Stdin, os.Stdout) }
