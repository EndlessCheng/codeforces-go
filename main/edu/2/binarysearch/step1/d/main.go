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

	var n, q, l, r int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &l, &r)
		Fprint(out, sort.SearchInts(a, r+1)-sort.SearchInts(a, l), " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }
