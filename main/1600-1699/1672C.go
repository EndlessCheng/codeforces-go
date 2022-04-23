package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1672C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, pre, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &pre)
		p := []int{}
		for i := 1; i < n; i++ {
			if Fscan(in, &v); v == pre {
				p = append(p, i)
			}
			pre = v
		}
		if len(p) < 2 {
			Fprintln(out, 0)
		} else {
			ans := p[len(p)-1] - p[0] - 1
			if ans < 1 {
				ans = 1
			}
			Fprintln(out, ans)
		}
	}
}

//func main() { CF1672C(os.Stdin, os.Stdout) }
