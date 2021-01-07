package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1470A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		id := make([]int, n)
		for i := range id {
			Fscan(in, &id[i])
		}
		sort.Ints(id)
		c := make([]int, m)
		for i := range c {
			Fscan(in, &c[i])
		}
		ans := int64(0)
		for i, k := range id {
			ans += int64(c[min(n-1-i, k-1)])
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1470A(os.Stdin, os.Stdout) }
