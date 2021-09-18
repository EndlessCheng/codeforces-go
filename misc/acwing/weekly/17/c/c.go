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
	var n, m, ans int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int, m)
	for i := range b {
		Fscan(in, &b[i])
	}
	for _, v := range a {
		i := sort.SearchInts(b, v)
		d := int(2e9)
		if i < m {
			d = b[i] - v
		}
		if i > 0 && v-b[i-1] < d {
			d = v - b[i-1]
		}
		if d > ans {
			ans = d
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
