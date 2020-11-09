package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF989C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct {
		v int
		c byte
	}

	a := make([]pair, 4)
	for i := byte(0); i < 4; i++ {
		Fscan(in, &a[i].v)
		a[i].c = 'A' + i
	}
	sort.Slice(a, func(i, j int) bool { return a[i].v < a[j].v })
	Fprintln(out, 50, 50)
	ans := [50][50]byte{}
	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			ans[i][j] = a[0].c
		}
	}
	for c, p := 0, a[3]; c < p.v; c++ {
		i, j := c/10*5+2, c%10*5+2
		ans[i-1][j] = p.c
		ans[i][j-1] = p.c
		if c+1 < a[0].v {
			ans[i-1][j-1] = p.c
		} else {
			ans[i][j] = p.c
		}
		if c < a[1].v {
			ans[i][j+1] = a[1].c
		}
		if c < a[2].v {
			ans[i+1][j] = a[2].c
		}
	}
	for _, row := range ans {
		Fprintf(out, "%s\n", row)
	}
}

//func main() { CF989C(os.Stdin, os.Stdout) }
