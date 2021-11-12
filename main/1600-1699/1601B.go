package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1601B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	dec := make([]int, n)
	for i := range dec {
		Fscan(in, &dec[i])
	}
	inc := make([]int, n)
	for i := range inc {
		Fscan(in, &inc[i])
	}

	vis := make([]bool, n)
	vis[n-1] = true
	type pair struct{ h, fa int }
	from := make([]pair, n)
	q := []int{n - 1}
	for step, min := 1, n-1; q != nil; step++ {
		tmp := q
		q = nil
		for _, v := range tmp {
			st := v - dec[v]
			if st < 0 {
				Fprintln(out, step)
				path := make([]int, 1, step)
				for ; v != n-1; v = from[v].fa {
					path = append(path, from[v].h+1)
				}
				for i := len(path) - 1; i >= 0; i-- {
					Fprint(out, path[i], " ")
				}
				return
			}
			if st >= min {
				continue
			}
			for h := st; h < min; h++ {
				if w := h + inc[h]; !vis[w] {
					vis[w] = true
					from[w] = pair{h, v}
					q = append(q, w)
				}
			}
			min = st
		}
	}
	Fprint(out, -1)
}

//func main() { CF1601B(os.Stdin, os.Stdout) }
