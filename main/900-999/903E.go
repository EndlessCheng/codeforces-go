package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF903E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m int
	Fscan(in, &n, &m)
	a := make([][]byte, n)
	cnt := make([][26]int, n)
	has2 := make([]bool, n)
	for i := range a {
		Fscan(in, &a[i])
		for _, b := range a[i] {
			if cnt[i][b-'a']++; cnt[i][b-'a'] > 1 {
				has2[i] = true
			}
		}
		if cnt[i] != cnt[0] {
			Fprint(out, -1)
			return
		}
	}

	s0 := a[0]
	dis := make([]int, n) // 汉明距离
	for i, s := range a {
		for j, b := range s {
			if b != s0[j] {
				dis[i]++
			}
		}
	}
	for l := 0; l < m; l++ {
	o:
		for r := l + 1; r < m; r++ {
			for i := 1; i < n; i++ {
				s, d := a[i], dis[i]
				if s[l] != s0[r] {
					d++
				}
				if s[r] != s0[l] {
					d++
				}
				if s[l] != s0[l] {
					d--
				}
				if s[r] != s0[r] {
					d--
				}
				if d > 2 || d == 0 && !has2[i] {
					continue o
				}
			}
			s0[l], s0[r] = s0[r], s0[l]
			Fprint(out, string(s0))
			return
		}
	}
	Fprint(out, -1)
}

//func main() { CF903E(os.Stdin, os.Stdout) }