package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func runG(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)

	ans := sort.Search(2e9, func(sum int) bool {
		c := 0
		var f func(p, s int)
		f = func(p, s int) {
			if p == n || c >= k || s+a[p] > sum {
				return
			}
			c++
			f(p+1, s+a[p])
			f(p+1, s)
		}
		f(0, 0)
		return c >= k
	})
	Fprint(out, ans)
}

//func main() { runG(os.Stdin, os.Stdout) }
