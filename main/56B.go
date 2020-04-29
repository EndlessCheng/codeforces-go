package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF56B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	var n, st, end int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if st == 0 && a[i] != i+1 {
			st, end = i+1, a[i]
		}
	}
	for i, j := st-1, end-1; i < j; i++ {
		a[i], a[j] = a[j], a[i]
		j--
	}
	for i, v := range a {
		if v != i+1 {
			st, end = 0, 0
			break
		}
	}
	Fprint(_w, st, end)
}

//func main() { CF56B(os.Stdin, os.Stdout) }
