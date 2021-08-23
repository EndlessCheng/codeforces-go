package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF549B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	g := make([][]byte, n)
	for i := range g {
		Fscan(in, &g[i])
	}
	q := []int{}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] == 0 {
			q = append(q, i)
		}
	}

	ans := []int{}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		ans = append(ans, v+1)
		for j, b := range g[v] {
			if b == '1' {
				if a[j]--; a[j] == 0 {
					q = append(q, j)
				}
			}
		}
	}
	sort.Ints(ans)
	Fprintln(out, len(ans))
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF549B(os.Stdin, os.Stdout) }
