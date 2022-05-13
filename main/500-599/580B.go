package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF580B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, d, l int
	Fscan(in, &n, &d)
	a := make([]struct{ x, y int }, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })
	var ans, s int64
	for _, p := range a {
		s += int64(p.y)
		for p.x-a[l].x >= d {
			s -= int64(a[l].y)
			l++
		}
		if s > ans {
			ans = s
		}
	}
	Fprint(out, ans)
}

//func main() { CF580B(os.Stdin, os.Stdout) }
