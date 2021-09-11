package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1419D2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	b := make([]int, n)
	j := -1
	for _, v := range a {
		if j += 2; j >= n {
			j = 0
		}
		b[j] = v
	}
	for i := 1; i < n-1; i++ {
		if b[i] < b[i-1] && b[i] < b[i+1] {
			ans++
		}
	}
	Fprintln(out, ans)
	for _, v := range b {
		Fprint(out, v, " ")
	}
}

//func main() { CF1419D2(os.Stdin, os.Stdout) }
