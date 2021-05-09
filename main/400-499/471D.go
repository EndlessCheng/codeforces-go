package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF471D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, p, v, ans int
	Fscan(in, &n, &m, &p)
	if m == 1 {
		Fprint(out, n)
		return
	}
	a := make([]int, n-1)
	for i := range a {
		Fscan(in, &v)
		a[i] = v - p
		p = v
	}
	Fscan(in, &p)
	b := make([]int, m-1)
	for i := range b {
		Fscan(in, &v)
		b[i] = v - p
		p = v
	}
	match := make([]int, m-1)
	for i, c := 1, 0; i < m-1; i++ {
		v := b[i]
		for c > 0 && b[c] != v {
			c = match[c-1]
		}
		if b[c] == v {
			c++
		}
		match[i] = c
	}
	c := 0
	for _, v := range a {
		for c > 0 && b[c] != v {
			c = match[c-1]
		}
		if b[c] == v {
			c++
		}
		if c == m-1 {
			ans++
			c = match[c-1]
		}
	}
	Fprint(out, ans)
}

//func main() { CF471D(os.Stdin, os.Stdout) }
