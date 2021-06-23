package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1261C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, m int
	Fscan(in, &n, &m)
	var s string
	a := make([][]int, n+1)
	b := make([][]int, n+2)
	a[0] = make([]int, m+1)
	b[n+1] = make([]int, m+2)
	for i := 0; i < n; i++ {
		a[i+1] = make([]int, m+1)
		b[i+1] = make([]int, m+2)
		Fscan(in, &s)
		for j, v := range s {
			if v == 'X' {
				a[i+1][j+1] = min(min(a[i][j], a[i][j+1]), a[i+1][j]) + 1
			}
		}
	}

	t := sort.Search(min(n, m), func(t int) bool {
		t++
		for i := n; i > 0; i-- {
			for j := m; j > 0; j-- {
				if a[i][j] > t*2 {
					b[i][j] = a[i][j]
				} else {
					b[i][j] = max(1, max(max(b[i+1][j], b[i+1][j+1]), b[i][j+1])) - 1
					if b[i][j] == 0 && a[i][j] > 0 {
						return true
					}
				}
			}
		}
		return false
	})
	Fprintln(out, t)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if i+t <= n && j+t <= m && a[i+t][j+t] > t*2 {
				Fprint(out, "X")
			} else {
				Fprint(out, ".")
			}
		}
		Fprintln(out)
	}
}

//func main() { CF1261C(os.Stdin, os.Stdout) }
