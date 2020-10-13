package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1370D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	f := func(c, up int) bool {
		for i := c; i < n; {
			for ; i < n && a[i] > up; i++ {
			}
			if i < n {
				c++
				i++
			}
			if i < n {
				c++
				i++
			}
		}
		return c >= k
	}
	Fprint(out, sort.Search(1e9, func(up int) bool { return f(0, up) || f(1, up) }))
}

//func main() { CF1370D(os.Stdin, os.Stdout) }
