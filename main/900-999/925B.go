package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF925B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, x1, x2 int
	Fscan(in, &n, &x1, &x2)
	a := make([]struct{ v, i int }, n)
	for i := range a {
		Fscan(in, &a[i].v)
		a[i].i = i + 1
	}
	sort.Slice(a, func(i, j int) bool { return a[i].v > a[j].v })

	swap := false
	f := func() bool {
		for i, p := range a {
			if p.v > (x1-1)/(i+1) {
				for j := i + 1; j < n; j++ {
					if a[j].v > (x2-1)/(j-i) {
						Fprintln(out, "Yes")
						if !swap {
							Fprintln(out, i+1, j-i)
							for _, p := range a[:i+1] {
								Fprint(out, p.i, " ")
							}
							Fprintln(out)
							for _, p := range a[i+1 : j+1] {
								Fprint(out, p.i, " ")
							}
						} else {
							Fprintln(out, j-i, i+1)
							for _, p := range a[i+1 : j+1] {
								Fprint(out, p.i, " ")
							}
							Fprintln(out)
							for _, p := range a[:i+1] {
								Fprint(out, p.i, " ")
							}
						}
						return true
					}
				}
				return false
			}
		}
		return false
	}
	if f() {
		return
	}
	x1, x2, swap = x2, x1, true
	if f() {
		return
	}
	Fprint(out, "No")
}

//func main() { CF925B(os.Stdin, os.Stdout) }
