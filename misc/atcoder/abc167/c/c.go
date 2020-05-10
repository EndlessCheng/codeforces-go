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
	var n, m, low int
	Fscan(in, &n, &m, &low)
	cost := make([]int, n)
	mat := make([][]int, n)
	for i := range mat {
		Fscan(in, &cost[i])
		mat[i] = make([]int, m)
		for j := range mat[i] {
			Fscan(in, &mat[i][j])
		}
	}

	ans := int(1e9)
	for sub := 0; sub < 1<<n; sub++ {
		s := 0
		level := make([]int, m)
		for i := 0; i < n; i++ {
			if sub>>i&1 == 1 {
				s += cost[i]
				for j, v := range mat[i] {
					level[j] += v
				}
			}
		}
		ok := true
		for _, v := range level {
			if v < low {
				ok = false
				break
			}
		}
		if ok && s < ans {
			ans = s
		}
	}
	if ans == 1e9 {
		ans = -1
	}
	Fprint(_w, ans)
}

func main() { run(os.Stdin, os.Stdout) }
