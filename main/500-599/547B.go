package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF547B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	type pair struct{ v, i int }
	posL := make([]int, n)
	s := []pair{{0, -1}}
	for i := range a {
		Fscan(in, &a[i])
		for {
			if top := s[len(s)-1]; top.v < a[i] {
				posL[i] = top.i
				break
			}
			s = s[:len(s)-1]
		}
		s = append(s, pair{a[i], i})
	}
	posR := make([]int, n)
	s = []pair{{0, n}}
	for i := n - 1; i >= 0; i-- {
		v := a[i]
		for {
			if top := s[len(s)-1]; top.v < v {
				posR[i] = top.i
				break
			}
			s = s[:len(s)-1]
		}
		s = append(s, pair{v, i})
	}

	ans := make([]int, n+1)
	for i, v := range a {
		sz := posR[i] - posL[i] - 1
		if v > ans[sz] {
			ans[sz] = v
		}
	}
	for i := n - 1; i > 0; i-- {
		if ans[i+1] > ans[i] {
			ans[i] = ans[i+1]
		}
	}
	for _, v := range ans[1:] {
		Fprint(out, v, " ")
	}
}

//func main() { CF547B(os.Stdin, os.Stdout) }
