package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1132C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	type pair struct{ x, y int }

	var n, m, l, r, all int
	Fscan(in, &n, &m)
	ids := make([][]int, n+1)
	for i := 0; i < m; i++ {
		Fscan(in, &l, &r)
		for p := l; p <= r; p++ {
			if len(ids[p]) < 3 {
				ids[p] = append(ids[p], i)
			}
		}
	}
	c1 := make([]int, m)
	c2 := map[pair]int{}
	for _, id := range ids {
		if id == nil {
			continue
		}
		if len(id) == 1 {
			c1[id[0]]++
		} else if len(id) == 2 {
			c2[pair{id[0], id[1]}]++
		}
		all++
	}
	min := n
	for i, c := range c1 {
		for j := i + 1; j < m; j++ {
			if c := c + c1[j] + c2[pair{i, j}]; c < min {
				min = c
			}
		}
	}
	Fprint(out, all-min)
}

//func main() { CF1132C(os.Stdin, os.Stdout) }
