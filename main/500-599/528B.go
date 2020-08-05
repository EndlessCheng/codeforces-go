package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF528B(_r io.Reader, out io.Writer) {
	type pair struct{ x, y int }
	in := bufio.NewReader(_r)
	var n, ans int
	Fscan(in, &n)
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y)
	}
	sort.Slice(a, func(i, j int) bool { x, y := a[i], a[j]; return x.x+x.y < y.x+y.y })
	cur := pair{-1e9, 0}
	for _, p := range a {
		if p.x-cur.x >= p.y+cur.y {
			ans++
			cur = p
		}
	}
	Fprint(out, ans)
}

//func main() { CF528B(os.Stdin, os.Stdout) }
