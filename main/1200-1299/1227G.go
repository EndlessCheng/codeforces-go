package main

import (
	"bytes"
	. "fmt"
	"io"
	"slices"
)

func cf1227G(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	type pair struct{ v, i int }
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].v)
		a[i].i = i
	}
	slices.SortFunc(a, func(a, b pair) int { return a.v - b.v })

	ans := make([][]byte, n+1)
	for i := range ans {
		ans[i] = bytes.Repeat([]byte{'0'}, n)
	}
	for i, p := range a {
		for j := 0; j < p.v; j++ {
			ans[(i+j)%(n+1)][p.i] = '1'
		}
	}
	Fprintln(out, n+1)
	for _, row := range ans {
		Fprintf(out, "%s\n", row)
	}
}

//func main() { cf1227G(bufio.NewReader(os.Stdin), os.Stdout) }
