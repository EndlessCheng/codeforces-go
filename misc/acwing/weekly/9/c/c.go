package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	type pair struct{ v, i int }
	posL := make([]int, n)
	s := []pair{{0, -1}}
	for i, v := range a {
		for {
			if top := s[len(s)-1]; top.v <= v {
				posL[i] = top.i
				break
			}
			s = s[:len(s)-1]
		}
		s = append(s, pair{v, i})
	}
	posR := make([]int, n)
	s = []pair{{0, n}}
	for i := n - 1; i >= 0; i-- {
		v := a[i]
		for {
			if top := s[len(s)-1]; top.v <= v {
				posR[i] = top.i
				break
			}
			s = s[:len(s)-1]
		}
		s = append(s, pair{v, i})
	}

	sl := make([]int, n)
	sl[0] = a[0]
	for i := 1; i < n; i++ {
		l := posL[i]
		sl[i] = a[i] * (i-l)
		if l >= 0 {
			sl[i] += sl[l]
		}
	}
	sr := make([]int, n)
	sr[n-1] = a[n-1]
	for i := n - 2; i >= 0; i-- {
		r := posR[i]
		sr[i] = a[i] * (r-i)
		if r < n {
			sr[i] += sr[r]
		}
	}

	maxS, maxI := 0, 0
	for i, v := range a {
		if sum := sl[i] + sr[i] - v; sum > maxS {
			maxS, maxI = sum, i
		}
	}

	ans := make([]int, n)
	curMax := a[maxI]
	for j := maxI; j >= 0; j-- {
		curMax = min(curMax, a[j])
		ans[j] = curMax
	}
	curMax = a[maxI]
	for j := maxI + 1; j < n; j++ {
		curMax = min(curMax, a[j])
		ans[j] = curMax
	}
	for _, v := range ans {
		Fprint(out, v, " ")
	}
	Fprintln(out)
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
