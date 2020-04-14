package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1098A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	var n, v int
	Fscan(in, &n)
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		Fscan(in, &v)
		v--
		g[v] = append(g[v], w)
	}
	s := make([]int, n)
	for i := range s {
		Fscan(in, &s[i])
	}

	ans := int64(0)
	var f func(int, int)
	f = func(v, lowS int) {
		sv := s[v]
		if sv != -1 && sv < lowS {
			ans = -1
			return
		}
		if len(g[v]) == 0 {
			if s[v] == -1 {
				s[v] = lowS
			}
			return
		}
		if sv > lowS {
			lowS = sv
		}
		minS := int(1e9)
		for _, w := range g[v] {
			f(w, lowS)
			if s[w] < minS {
				minS = s[w]
			}
		}
		if s[v] == -1 {
			s[v] = minS
		}
	}
	f(0, s[0])
	a := make([]int, n)
	f = func(v, sFa int) {
		a[v] = s[v] - sFa
		for _, w := range g[v] {
			f(w, s[v])
		}
	}
	f(0, 0)
	if ans != -1 {
		for _, v := range a {
			ans += int64(v)
		}
	}
	Fprint(_w, ans)
}

//func main() { CF1098A(os.Stdin, os.Stdout) }
