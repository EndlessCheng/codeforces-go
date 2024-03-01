package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf845C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]struct{ l, r int }, n)
	for i := range a {
		Fscan(in, &a[i].l, &a[i].r)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].l < a[j].l })

	end1, end2 := -1, -1
	for _, p := range a {
		if p.l > end1 {
			end1 = p.r
		} else if p.l > end2 {
			end2 = p.r
		} else {
			Fprint(out, "NO")
			return
		}
	}
	Fprint(out, "YES")
}

//func main() { cf845C(os.Stdin, os.Stdout) }
