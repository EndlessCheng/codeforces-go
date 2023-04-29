package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1479B2(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n ,ans int
	Fscan(in, &n)
	a := make([]int, n)
	pos := make([][]int, n+1)
	for i := range a {
		Fscan(in, &a[i])
		pos[a[i]] = append(pos[a[i]], i)
	}
	for i := range pos {
		pos[i] = append(pos[i], n)
	}
	s, t := []int{0}, []int{0}
	for _, v := range a {
		if v == s[len(s)-1] {
			s = append(s, v)
		} else if v == t[len(t)-1] {
			t = append(t, v)
		} else if pos[s[len(s)-1]][0] > pos[t[len(t)-1]][0] {
			s = append(s, v)
		} else {
			t = append(t, v)
		}
		pos[v] = pos[v][1:]
	}
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			ans++
		}
	}
	for i := 1; i < len(t); i++ {
		if t[i] != t[i-1] {
			ans++
		}
	}
	Fprint(out, ans)
}

//func main() { CF1479B2(os.Stdin, os.Stdout) }
