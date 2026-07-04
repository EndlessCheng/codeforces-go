package main

import (
	"bufio"
	"cmp"
	. "fmt"
	"io"
	"os"
	"slices"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, c1 int
	var s []byte
	Fscan(in, &n, &s)
	type pair struct {
		v int
		c int8
	}
	a := make([]pair, n)
	for i, c := range s {
		Fscan(in, &a[i].v)
		a[i].c = int8(c)
		c1 += int(c - '0')
	}
	slices.SortFunc(a, func(a, b pair) int { return cmp.Or(a.v-b.v, int(b.c-a.c)) })

	ans, f := c1, c1
	for _, p := range a {
		if p.c == '0' {
			f++
			ans = max(ans, f)
		} else {
			f--
		}
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
