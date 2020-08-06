package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k int64
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
	}
	sort.Ints(b)
	Fprint(out, sort.Search(2e9, func(s int) bool {
		c := int64(0)
		for _, v := range a {
			c += int64(sort.SearchInts(b, s-v+1))
		}
		return c >= k
	}))
}

func main() { run(os.Stdin, os.Stdout) }
