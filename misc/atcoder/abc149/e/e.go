package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, ans int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	minSum := sort.Search(2e5+1, func(s int) bool {
		c := 0
		for i := n - 1; i >= 0; i-- {
			j := sort.SearchInts(a, s-a[i])
			c += n - j
		}
		return c < m
	}) - 1
	s := make([]int, n+1)
	for i, v := range a {
		s[i+1] = s[i] + v
	}
	for _, v := range a {
		j := sort.SearchInts(a, minSum-v+1)
		ans += s[n] - s[j] + (n-j)*v
		m -= n - j
	}
	ans += m * minSum
	Fprint(_w, ans)
}

func main() { run(os.Stdin, os.Stdout) }
