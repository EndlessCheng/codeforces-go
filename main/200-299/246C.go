package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF246C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	for m := n; ; m-- {
		for _, v := range a[:m] {
			Fprint(out, n-m+1, v, " ")
			for _, w := range a[m:] {
				Fprint(out, w, " ")
			}
			Fprintln(out)
			if k--; k == 0 {
				return
			}
		}
	}
}

//func main() { CF246C(os.Stdin, os.Stdout) }
