package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1329A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, m int
	Fscan(in, &n, &m)
	l := make([]int, m)
	r := make([]int, m)
	p := make([]int, m)
	for i := range l {
		Fscan(in, &l[i])
		if i == 0 {
			r[i] = l[i]
		} else {
			r[i] = max(r[i-1], l[i]+i)
		}
		if r[i] > n {
			Fprint(out, -1)
			return
		}
		p[i] = i + 1
	}
	last := n
	i := m - 1
	for ; i >= 0 && r[i] < last; i-- {
		p[i] = last - l[i] + 1
		last = p[i] - 1
	}
	if i < 0 && last > 0 {
		Fprint(out, -1)
		return
	}
	for _, v := range p {
		Fprint(out, v, " ")
	}
}

//func main() { CF1329A(os.Stdin, os.Stdout) }
