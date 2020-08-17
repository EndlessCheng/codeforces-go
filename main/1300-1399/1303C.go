package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1303C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func() (ans []byte) {
		var s []byte
		Fscan(in, &s)

		used := [26]bool{}
		if len(s) > 1 {
			for i := range s {
				s[i] -= 'a'
			}
			g := [26]map[byte]int{}
			for i := range g {
				g[i] = map[byte]int{}
			}
			for i := 0; i < len(s)-1; i++ {
				v, w := s[i], s[i+1]
				g[v][w]++
				g[w][v]++
			}
			st := -1
			for v, e := range g {
				if len(e) == 1 {
					st = v
				} else if len(e) > 2 {
					return
				}
			}
			if st == -1 {
				return
			}
			var f func(byte)
			f = func(v byte) {
				used[v] = true
				ans = append(ans, 'a'+v)
				for w := range g[v] {
					if !used[w] {
						f(w)
					}
				}
			}
			f(byte(st))
		}
		for i, u := range used {
			if !u {
				ans = append(ans, byte('a'+i))
			}
		}
		return
	}

	var t int
	for Fscan(in, &t); t > 0; t-- {
		if ans := solve(); ans != nil {
			Fprintf(out, "YES\n%s\n", ans)
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1303C(os.Stdin, os.Stdout) }
