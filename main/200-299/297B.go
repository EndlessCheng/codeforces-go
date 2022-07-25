package main

import (
	"bufio"
	. "fmt"
	"io"
	. "sort"
)

// github.com/EndlessCheng/codeforces-go
func CF297B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, k int
	Fscan(in, &n, &m, &k)
	if n > m {
		Fprint(out, "YES")
		return
	}
	a := make(IntSlice, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	Sort(Reverse(a))
	b := make(IntSlice, m)
	for i := range b {
		Fscan(in, &b[i])
	}
	Sort(Reverse(b))

	for i, v := range a {
		if v > b[i] {
			Fprint(out, "YES")
			return
		}
	}
	Fprint(out, "NO")
}

//func main() { CF297B(os.Stdin, os.Stdout) }
