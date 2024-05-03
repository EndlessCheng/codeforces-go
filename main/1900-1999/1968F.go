package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1968F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, q, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		s := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &s[i])
			s[i] ^= s[i-1]
		}
		pos := map[int][]int{}
		for i, v := range s {
			pos[v] = append(pos[v], i)
		}
		for ; q > 0; q-- {
			Fscan(in, &l, &r)
			l--
			if s[l] == s[r] {
				Fprintln(out, "YES")
				continue
			}
			p, q := pos[s[r]], pos[s[l]]
			if p[sort.SearchInts(p, l)] < q[sort.SearchInts(q, r)-1] {
				Fprintln(out, "YES")
			} else {
				Fprintln(out, "NO")
			}
		}
	}
}

//func main() { cf1968F(os.Stdin, os.Stdout) }
