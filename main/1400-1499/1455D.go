package main

import (
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1455D(in io.Reader, out io.Writer) {
	var T, n, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		p, ans := 0, 0
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			if i > 0 && a[i-1] > a[i] {
				p = i
			}
		}
		for i, v := range a[:p] {
			if x < v {
				x, a[i] = v, x
				ans++
			}
		}
		if sort.IntsAreSorted(a) {
			Fprintln(out, ans)
		} else {
			Fprintln(out, -1)
		}
	}
}

//func main() { CF1455D(os.Stdin, os.Stdout) }
