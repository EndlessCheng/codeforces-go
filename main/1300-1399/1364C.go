package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1364C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, cur int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int, n)
	for i := range b {
		b[i] = -1
	}
	used := make([]bool, n+1)
	for v := 1; ; {
		i := sort.SearchInts(a, v)
		if i == n {
			break
		}
		if i == 0 {
			b[i] = 0
		} else {
			b[i] = a[i-1]
		}
		used[b[i]] = true
		v = a[i] + 1
	}
	used[a[n-1]] = true
	for _, v := range b {
		if v < 0 {
			for ; used[cur]; cur++ {
			}
			v = cur
			cur++
		}
		Fprint(out, v, " ")
	}
}

//func main() { CF1364C(os.Stdin, os.Stdout) }
