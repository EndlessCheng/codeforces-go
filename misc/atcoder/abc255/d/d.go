package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, x int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	sum := make([]int, n+1)
	for i, x := range a {
		sum[i+1] = sum[i] + x
	}
	for ; q > 0; q-- {
		Fscan(in, &x)
		j := sort.SearchInts(a, x)
		Fprintln(out, x*j+sum[n]-sum[j]*2-x*(n-j))
	}
}

func main() { run(os.Stdin, os.Stdout) }
