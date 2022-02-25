package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1644C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var T, n, x, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		s := make([]int, n+1)
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			s[i+1] = s[i] + v
		}
		mx := make([]int, n+1)
		mx[n] = s[n]
		for l := n - 1; l >= 0; l-- {
			mx[l] = mx[l+1]
			for i := l; i <= n; i++ {
				mx[l] = max(mx[l], s[i]-s[i-l])
			}
		}
		ans := 0
		for i, v := range mx {
			ans = max(ans, v+i*x)
			Fprint(out, ans, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1644C(os.Stdin, os.Stdout) }
