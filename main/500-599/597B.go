package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF597B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, pre, ans int
	Fscan(in, &n)
	a := make([]struct{ l, r int }, n)
	for i := range a {
		Fscan(in, &a[i].l, &a[i].r)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].r < a[j].r })
	for _, p := range a {
		if p.l > pre {
			ans++
			pre = p.r
		}
	}
	Fprint(out, ans)
}

//func main() { CF597B(os.Stdin, os.Stdout) }
