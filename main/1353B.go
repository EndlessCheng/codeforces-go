package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1353B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, k int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := make([]int, n)
		for i := range b {
			Fscan(in, &b[i])
		}
		sort.Ints(b)
		a = append(a, b[n-k:]...)
		sort.Ints(a)
		s := 0
		for _, v := range a[k:] {
			s += v
		}
		Fprintln(out, s)
	}
}

//func main() { CF1353B(os.Stdin, os.Stdout) }
