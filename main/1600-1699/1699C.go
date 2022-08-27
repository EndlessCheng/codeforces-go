package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1699C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod int64 = 1e9 + 7

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		p := make([]int, n)
		for i := range p {
			Fscan(in, &v)
			p[v] = i
		}
		ans := int64(1)
		l, r := p[0], p[0]
		for i, p := range p[1:] {
			if p < l {
				l = p
			} else if p > r {
				r = p
			} else {
				ans = ans * int64(r-l-i) % mod
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1699C(os.Stdin, os.Stdout) }
