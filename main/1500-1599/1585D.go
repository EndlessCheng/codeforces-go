package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1585D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n+1)
		p := make([]int, n+1)
		eq := false
		for i := 1; i <= n; i++ { // 注意 a 从 1 开始
			Fscan(in, &a[i])
			if p[a[i]] > 0 {
				eq = true
			}
			p[a[i]] = i
		}
		if eq || n < 2 {
			Fprintln(out, "YES")
			continue
		}
		for i := 1; i < n-1; i++ {
			if p[i] == i {
				continue
			}
			x, y := p[i], i
			z := n
			if x == n {
				z--
			}
			p[a[z]] = x
			p[a[y]] = z
			a[x] = a[z]
			a[z] = a[y]
			//a[x], a[y], a[z] = a[z], a[x], a[y]
		}
		if a[n-1] < a[n] { // 注意 a 从 1 开始
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1585D(os.Stdin, os.Stdout) }
