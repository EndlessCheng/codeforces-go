package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	type pair struct{ x, y int }
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].x)
	}
	for i := range a {
		Fscan(in, &a[i].y)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })
	f := []int{}
	for _, p := range a {
		v := p.y
		j := sort.SearchInts(f, v)
		if j < len(f) {
			f[j] = v
		} else {
			f = append(f, v)
		}
	}
	Fprint(out, n+len(f))
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
