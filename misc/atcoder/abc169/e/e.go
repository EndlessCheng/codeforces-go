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

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	b := make([]int, n)
	for i := range a {
		Fscan(in, &a[i], &b[i])
	}
	sort.Ints(a)
	sort.Ints(b)
	min, max := a[n/2], b[n/2]
	if n&1 == 0 {
		min += a[n/2-1]
		max += b[n/2-1]
	}
	Fprint(out, max-min+1)
}

func main() { run(os.Stdin, os.Stdout) }
