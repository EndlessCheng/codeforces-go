package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1324D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	var n, v int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for i := range a {
		Fscan(in, &v)
		a[i] -= v
	}
	sort.Ints(a)
	ans := int64(0)
	for i, v := range a {
		if j := sort.SearchInts(a[:i], 1-v); j < i {
			ans += int64(i - j)
		}
	}
	Fprint(_w, ans)
}

//func main() { CF1324D(os.Stdin, os.Stdout) }
