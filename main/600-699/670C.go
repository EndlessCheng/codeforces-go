package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF670C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	c := map[int]int{}
	var n, m, w, ans, c1, c2 int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &w)
		c[w]++
	}
	Fscan(in, &m)
	a := make([]int, m)
	for i := range a {
		Fscan(in, &a[i])
	}
	for i, v := range a {
		if Fscan(in, &w); c[v] > c1 {
			ans, c1, c2 = i, c[v], c[w]
		} else if c[v] == c1 && c[w] > c2 {
			ans, c2 = i, c[w]
		}
	}
	Fprint(out, ans+1)
}

//func main() { CF670C(os.Stdin, os.Stdout) }
