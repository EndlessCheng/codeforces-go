package main

import (
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF626C(in io.Reader, out io.Writer) {
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n, m int
	Fscan(in, &n, &m)
	Fprint(out, sort.Search(3*(n+m), func(lim int) bool {
		return (max(n-(lim+4)/6-(lim+2)/6, 0)+max(m-(lim+3)/6, 0))*6 <= lim
	}))
}

//func main() { CF626C(os.Stdin, os.Stdout) }
