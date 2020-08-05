package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF294B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	b := [3][]int{}
	var n, t, w, all, s int
	Fscan(in, &n)
	for i := 0; i < n; i++ {
		Fscan(in, &t, &w)
		b[t] = append(b[t], w)
		all += w
	}
	for i := 1; i < 3; i++ {
		sort.Sort(sort.Reverse(sort.IntSlice(b[i])))
		b[i] = append(b[i], 0)
	}
	ans := int(1e9)
	for i, v := range b[1] {
		ss := s
		for j, w := range b[2] {
			if t := i + j*2; t >= all-ss && t < ans {
				ans = t
			}
			ss += w
		}
		s += v
	}
	Fprint(out, ans)
}

//func main() { CF294B(os.Stdin, os.Stdout) }
