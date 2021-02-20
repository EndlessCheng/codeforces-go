package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1490G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, q int
	var x int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		s := make([]int64, n+1)
		p := []int{0}
		for i := 1; i <= n; i++ {
			Fscan(in, &x)
			s[i] = s[i-1] + x
			if s[i] > s[p[len(p)-1]] {
				p = append(p, i)
			}
		}
		for ; q > 0; q-- {
			Fscan(in, &x)
			if x <= s[p[len(p)-1]] {
				i := sort.Search(len(p), func(i int) bool { return s[p[i]] >= x })
				Fprint(out, p[i]-1, " ")
			} else if s[n] < 1 {
				Fprint(out, "-1 ")
			} else {
				loop := (x-s[p[len(p)-1]]-1)/s[n] + 1
				x -= loop * s[n]
				i := sort.Search(len(p), func(i int) bool { return s[p[i]] >= x })
				Fprint(out, loop*int64(n)+int64(p[i]-1), " ")
			}
		}
		Fprintln(out)
	}
}

//func main() { CF1490G(os.Stdin, os.Stdout) }
