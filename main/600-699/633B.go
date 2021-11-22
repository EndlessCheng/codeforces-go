package main

import (
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF633B(in io.Reader, out io.Writer) {
	var m int
	Fscan(in, &m)
	ok := false
	n := sort.Search(1e6, func(n int) bool {
		k := 0
		for n > 0 {
			n /= 5
			k += n
		}
		ok = ok || k == m
		return k >= m
	})
	if ok {
		Fprintln(out, 5)
		Fprint(out, n, n+1, n+2, n+3, n+4)
	} else {
		Fprint(out, 0)
	}
}

//func main() { CF633B(os.Stdin, os.Stdout) }
