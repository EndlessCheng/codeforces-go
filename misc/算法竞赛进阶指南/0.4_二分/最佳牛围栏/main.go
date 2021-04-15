package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://www.luogu.com.cn/problem/P1404
// O(n) https://www.luogu.com.cn/blog/user43145/solution-p1404

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		a[i] *= 1e3
	}
	Fprint(out, sort.Search(2e6+1, func(avg int) bool {
		s := make([]int, n+1)
		for i, v := range a {
			s[i+1] = s[i] + v - avg
		}
		mi := 0
		for i, v := range s[:n+1-m] {
			if v < mi {
				mi = v
			}
			if s[i+m] >= mi {
				return false
			}
		}
		return true
	})-1)
}

func main() { run(os.Stdin, os.Stdout) }
